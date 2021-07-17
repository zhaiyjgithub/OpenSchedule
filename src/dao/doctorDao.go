package dao

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

type DoctorDao struct {
	elasticSearchEngine *elastic.Client
}

func NewDoctorDao(engine *elastic.Client) *DoctorDao  {
	return &DoctorDao{elasticSearchEngine: engine}
}

/*
|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|keyword |yes  |string |   |
|IsInClinicBookEnable |no  |bool |  default true  |
|IsVirtualBookEnable     |no  |bool |    |
|appointmentType | yes | int | all = 0 inClinic=1 virtual = 2
|nextAvailableDate |no |date | UTC format
|city |no |string |
|specialty |no | string |
|lat | yes | float |
|lon |yes | float |
|gender |no | bool|
|page |yes | int| default index from 1
|pageSize |yes | int | default size from 50*/
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
	pageSize int)  {

	q := elastic.NewBoolQuery()

	if len(keyword) > 0 {
		fuzzyQuery := elastic.NewFuzzyQuery("FullName", keyword).Boost(1.5).Fuzziness(2).PrefixLength(0).MaxExpansions(100)
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

	distanceQuery := elastic.NewGeoDistanceQuery("Location").Lat(lat).Lon(lon).Distance("200km").DistanceType("plane")
	q.Filter(distanceQuery)

	src, err := q.Source()
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	fmt.Println("got: ", got)
	sorter := elastic.NewGeoDistanceSort("Location").
		Point(lat, lon).
		Order(true).
		Unit("km").
		GeoDistance("plane")

	result, err := d.elasticSearchEngine.Search().Index(database.DoctorIndexName).
		Size(pageSize).
		From((page -1)*pageSize).
		Query(q).Pretty(true).SortBy(sorter).
		Do(context.Background())

	var doctors []*model.Doctor
	for _, hit := range result.Hits.Hits {
		var doc model.Doctor
		err = json.Unmarshal(hit.Source, &doc)

		if err != nil {
			continue
		}

		doctors = append(doctors, &doc)
	}

	//fmt.Println("doctors:", doctors)
	for _, doc := range doctors {
		fmt.Println(doc.FullName)
	}
}
