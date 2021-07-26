package dao
import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/viewModel"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

type ScriptLocation struct {
	Lat float64
	Lon float64
}

type DoctorDao struct {
	elasticSearchEngine *elastic.Client
}

func NewDoctorDao(engine *elastic.Client) *DoctorDao  {
	return &DoctorDao{elasticSearchEngine: engine}
}

func (d *DoctorDao) SearchDoctor(keyword string,
	IsInClinicBookEnable bool,
	IsVirtualBookEnable bool,
	appointmentType constant.AppointmentType,
	nextAvailableDate string,
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int, sortType constant.SortType, distance int ) []*viewModel.DoctorInfo {
	q := elastic.NewBoolQuery()
	if len(keyword) > 0 {
		fuzzyQuery := elastic.NewMatchQuery("FullName", keyword).Boost(1.5).Fuzziness("2").PrefixLength(0).MaxExpansions(100)
		q.Must(fuzzyQuery)
		q.Filter(elastic.NewTermQuery("IsInClinicBookEnable", IsInClinicBookEnable))
		q.Filter(elastic.NewTermQuery("IsVirtualBookEnable", IsVirtualBookEnable))
		if len(city) > 0 {
			q.Filter(elastic.NewTermQuery("City", city))
		}
		if len(specialty) > 0 {
			q.Filter(elastic.NewTermQuery("Specialty", specialty))
		}
		q.Filter(elastic.NewTermQuery("Gender", gender))
		if len(nextAvailableDate) > 0 {
			if appointmentType == constant.InClinic {
				q.Filter(elastic.NewRangeQuery("NextAvailableDateInClinic").
					Gte(nextAvailableDate))
			}else {
				q.Filter(elastic.NewRangeQuery("NextAvailableDateVirtual").
					Gte(nextAvailableDate))
			}
		}
	}else {
		q.Must(elastic.NewTermQuery("IsInClinicBookEnable", IsInClinicBookEnable))
		q.Must(elastic.NewTermQuery("IsVirtualBookEnable", IsVirtualBookEnable))
		if len(city) > 0 {
			q.Must(elastic.NewTermQuery("City", city))
		}
		if len(specialty) > 0 {
			q.Must(elastic.NewTermQuery("Specialty", specialty).Boost(3))
		}
		q.Filter(elastic.NewTermQuery("Gender", gender))
		if len(nextAvailableDate) > 0 {
			if appointmentType == constant.InClinic {
				q.Filter(elastic.NewRangeQuery("NextAvailableDateInClinic").
					Gte(nextAvailableDate))
			}else {
				q.Filter(elastic.NewRangeQuery("NextAvailableDateVirtual").
					Gte(nextAvailableDate))
			}
		}
	}

	defaultDistance := 10 //mini = 10km
	if distance > defaultDistance {
		defaultDistance = distance
	}
	distanceRange := fmt.Sprintf("%dkm", distance)
	distanceQuery := elastic.NewGeoDistanceQuery("Location").Lat(lat).Lon(lon).Distance(distanceRange).DistanceType("plane")
	q.Filter(distanceQuery)
	docs := make([]*viewModel.DoctorInfo, 0)
	if sortType == constant.ByDistance {
		docs = d.searchByDistance(lat, lon, q, page, pageSize)
	}else {
		docs = d.searchByDefault(lat, lon, q, page, pageSize)
	}
	return docs
}

func (d *DoctorDao)searchByDistance(lat float64, lon float64, q elastic.Query, page int , pageSize int) []*viewModel.DoctorInfo {
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
		return docs
	}
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

	return docs
}

func (d *DoctorDao)searchByDefault(lat float64, lon float64, q elastic.Query, page int , pageSize int) []*viewModel.DoctorInfo {
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
		return docs
	}
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
	return docs
}

//func sourceSQL(src, err)  {
//	src, err := builder.Source()
//	if err != nil {
//		log.Fatal(err)
//	}
//	data, err := json.Marshal(src)
//	if err != nil {
//		log.Fatalf("marshaling to JSON failed: %v", err)
//	}
//	got := string(data)
//	fmt.Println("got elastic search sql: ", got)
//}