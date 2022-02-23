package doctor

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/model/viewModel"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type ScriptLocation struct {
	Lat float64
	Lon float64
}

type Hits struct {
	TotalHits int64
}

type Dao struct {
	elasticSearchEngine *elastic.Client
	mainEngine *gorm.DB
}

func NewDoctorDao(engine *elastic.Client, mainEngine *gorm.DB) *Dao {
	return &Dao{elasticSearchEngine: engine, mainEngine: mainEngine}
}

func (d *Dao) SearchDoctor(keyword string,
	appointmentType constant.AppointmentType,
	startDate interface{},
	endDate interface{},
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int,
	sortType constant.SortByType,
	distance int,
	) (int64, []*viewModel.DoctorInfo) {
	q := elastic.NewBoolQuery()

	isInClinicEnable:= true
	isVirtualEnable:= true
	if appointmentType == constant.InClinic {
		isInClinicEnable = true
		isVirtualEnable = false
	} else if appointmentType == constant.Virtual {
		isInClinicEnable = false
		isVirtualEnable = true
	}
	if len(keyword) > 0 {
		fuzzyQuery := elastic.NewMatchQuery("FullName", keyword).Boost(1.5).Fuzziness("2").PrefixLength(0).MaxExpansions(100)
		q.Must(fuzzyQuery)
		if appointmentType == constant.InClinic {
			q.Filter(elastic.NewTermQuery("IsInClinicBookEnable", isInClinicEnable))
		} else if appointmentType == constant.Virtual {
			q.Filter(elastic.NewTermQuery("IsVirtualBookEnable", isVirtualEnable))
		} else {
			//
		}

		if len(city) > 0 {
			q.Filter(elastic.NewTermQuery("City", city))
		}
		if len(specialty) > 0 {
			q.Filter(elastic.NewTermQuery("Specialty", specialty))
		}
		if gender != constant.Trans {
			q.Filter(elastic.NewTermQuery("Gender", gender))
		}
		q = AddDateRangeQuery(q, appointmentType, startDate, endDate)
	}else {
		if appointmentType == constant.InClinic {
			q.Filter(elastic.NewTermQuery("IsInClinicBookEnable", isInClinicEnable))
		} else if appointmentType == constant.Virtual {
			q.Filter(elastic.NewTermQuery("IsVirtualBookEnable", isVirtualEnable))
		} else {
			//
		}
		if len(city) > 0 {
			q.Must(elastic.NewTermQuery("City", city))
		}
		if len(specialty) > 0 {
			q.Must(elastic.NewTermQuery("Specialty", specialty).Boost(3))
		}
		if gender != constant.Trans {
			q.Filter(elastic.NewTermQuery("Gender", gender))
		}
		q = AddDateRangeQuery(q, appointmentType, startDate, endDate)
	}

	defaultDistance := 1000 //default radius = 1000km for near by
	if distance != defaultDistance {
		defaultDistance = distance
	}
	distanceRange := fmt.Sprintf("%dkm", distance)
	distanceQuery := elastic.NewGeoDistanceQuery("Location").Lat(lat).Lon(lon).Distance(distanceRange).DistanceType("plane")
	q.Filter(distanceQuery)
	docs := make([]*viewModel.DoctorInfo, 0)
	total := int64(0)
	if sortType == constant.ByDistance {
		total, docs = d.searchByDistance(lat, lon, q, page, pageSize)
	}else {
		total, docs = d.searchByDefault(lat, lon, q, page, pageSize)
	}
	return total, docs
}

func AddDateRangeQuery(q *elastic.BoolQuery, appointmentType constant.AppointmentType, startDate interface{}, endDate interface{}) *elastic.BoolQuery  {
	if appointmentType == constant.InClinic {
		q.Should(elastic.NewRangeQuery("NextAvailableDateInClinic").
			Gte(startDate).Lte(endDate))
	}else if appointmentType == constant.Virtual {
		q.Should(elastic.NewRangeQuery("NextAvailableDateVirtual").
			Gte(startDate).Lte(endDate))
	} else {
		q.Should(elastic.NewRangeQuery("NextAvailableDateInClinic").
			Gte(startDate).Lte(endDate))
		q.Should(elastic.NewRangeQuery("NextAvailableDateVirtual").
			Gte(startDate).Lte(endDate))
	}
	return q
}

func (d *Dao)searchByDistance(lat float64, lon float64, q elastic.Query, page int , pageSize int) (int64, []*viewModel.DoctorInfo) {
	src, err := q.Source()
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	fmt.Println("got elastic search sql: ", got)

	sorter := elastic.NewGeoDistanceSort("Location").
		Point(lat, lon).
		Order(true).
		Unit("km").
		GeoDistance("plane")
	docs := make([]*viewModel.DoctorInfo, 0)
	result, err := d.elasticSearchEngine.Search().Index(database.DoctorIndexName).
		Size(pageSize).
		From((page -1)*pageSize).
		Query(q).Pretty(true).SortBy(sorter).
		Do(context.Background())
	if err != nil {
		fmt.Println("search failed")
		return 0, docs
	}

	totalHits := result.TotalHits()
	for _, hit := range result.Hits.Hits {
		var doc viewModel.DoctorInfo
		err = json.Unmarshal(hit.Source, &doc)
		if err != nil {
			continue
		}
		for _, val := range hit.Sort {
			distance, ok := val.(float64)
			if ok {
				doc.Distance = distance
			}
		}

		docs = append(docs, &doc)
	}
	for _, doc := range docs {
		fmt.Println(doc.FullName)
	}

	return totalHits, docs
}

func (d *Dao)searchByDefault(lat float64, lon float64, q elastic.Query, page int , pageSize int) (int64, []*viewModel.DoctorInfo) {
	sl := &ScriptLocation{
		Lat: lat,
		Lon: lon,
	}
	script := elastic.NewScript("doc['Location'].arcDistance(params.location.Lat,params.location.Lon)/1000").Param("location", sl).Lang("painless")
	sf := elastic.NewScriptField("distance", script)
	builder := elastic.NewSearchSource().Query(q).ScriptFields(sf).FetchSource(true).From((page - 1)*pageSize).Size(pageSize)
	src, err := builder.Source()
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	fmt.Println("got elastic search sql: ", got)

	docs := make([]*viewModel.DoctorInfo, 0)
	result, err := d.elasticSearchEngine.Search().Index(database.DoctorIndexName).SearchSource(builder).Do(context.Background())
	if err != nil {
		fmt.Println("search failed")
		return 0, docs
	}

	totalHits := result.TotalHits()
	for _, hit := range result.Hits.Hits {
		var doc viewModel.DoctorInfo
		err = json.Unmarshal(hit.Source, &doc)
		if err != nil {
			continue
		}
		if hit.Fields != nil {
			field, _ := hit.Fields["distance"]
			fields, ok := field.([]interface{})
			if !ok {
				fmt.Printf("expected []interface{}; got: %v\n", reflect.TypeOf(fields))
			}
			distance, ok := fields[0].(float64)
			if !ok {
				fmt.Printf("expected a string; got: %v\n", reflect.TypeOf(fields[0]))
			}
			doc.Distance = distance
		}
		docs = append(docs, &doc)
	}
	for _, doc := range docs {
		fmt.Println(doc.FullName)
	}
	return totalHits, docs
}

func (d *Dao) GetDoctorTimeSlots(npi int64)  {

}

func (d *Dao) GetDoctorByPage(page int , pageSize int) []*doctor.Doctor {
	var doctors []*doctor.Doctor
	_ = d.mainEngine.Limit(pageSize).Offset(pageSize*(page - 1)).Find(&doctors)
	return doctors
}

func (d *Dao) GetDoctor(npi int64) doctor.Doctor {
	var doc doctor.Doctor
	_ = d.mainEngine.Where("npi = ?", npi).First(&doc)
	return doc
}

func (d *Dao) IsExist(npi int64) bool {
	var count int64
	_ = d.mainEngine.Model(&doctor.Doctor{}).Where("npi = ?",  npi).Count(&count)
	return count > 0
}

func (d *Dao) SaveDoctor(doc *doctor.Doctor) error {
	var doctor doctor.Doctor
	db := d.mainEngine.Where("npi = ?", doc.Npi).First(&doctor)
	if db.Error != nil {
		return db.Error
	}
	doc.ID = doctor.ID
	db = d.mainEngine.Save(doc)
	return db.Error
}