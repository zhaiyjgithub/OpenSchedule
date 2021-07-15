package dao

import (
	"OpenSchedule/src/constant"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
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
|isInClinicEnable |no  |bool |  default true  |
|isVirtualEnable     |no  |bool |    |
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
	isInClinicEnable bool,
	isVirtualEnable bool,
	appointmentType constant.AppointmentType,
	nextAvailableDate time.Time,
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int)  {

	q := elastic.NewBoolQuery()

	if len(keyword) > 0 {
		fuzzyQuery := elastic.NewFuzzyQuery("FullName", keyword).Boost(1.5).Fuzziness(2).PrefixLength(0).MaxExpansions(0)
		q.Must(fuzzyQuery)

		q.Filter(elastic.NewTermQuery("IsInClinicEnable", isInClinicEnable))
		q.Filter(elastic.NewTermQuery("IsVirtualEnable", isVirtualEnable))
		q.Filter(elastic.NewTermQuery("AppointmentType", appointmentType))
		q.Filter(elastic.NewTermQuery("City", city))
		q.Filter(elastic.NewTermQuery("Specialty", specialty))
		q.Filter(elastic.NewTermQuery("Gender", gender))
		q.Filter(elastic.NewRangeQuery("born").
			Gte("2012-01-01"))

	}else {
		q.Must(elastic.NewTermQuery("IsInClinicEnable", isInClinicEnable))
		q.Must(elastic.NewTermQuery("IsVirtualEnable", isVirtualEnable))
		q.Must(elastic.NewTermQuery("AppointmentType", appointmentType))
		q.Must(elastic.NewTermQuery("City", city))
		q.Must(elastic.NewTermQuery("Specialty", specialty))
		q.Must(elastic.NewTermQuery("Gender", gender))
		q.Filter(elastic.NewRangeQuery("born").
			Gte("2012-01-01")).Boost(3)
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
	//sort := elastic.NewGeoDistanceSort("Location").
	//	Point(lat, lon).
	//	Order(true).
	//	Unit("km").
	//	SortMode("min").
	//	GeoDistance("plane")

}
