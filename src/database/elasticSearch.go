package database

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

type elasticIndex struct {
	Name string
	Mapping string
}

const (
	DoctorIndexName = "doctor"
	DoctorIndexMapping = `
	{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties": {
			  "Npi":    { "type": "long" },  
			  "FirstName":  { "type": "keyword"}, 
			  "LastName":   { "type": "keyword"},
			  "MiddleName": { "type": "keyword"},
			  "FullName": {"type": "text"},
			  "NamePrefix":  { "type": "keyword"}, 
			  "JobTitle":   { "type": "keyword"},
			  "Gender": {"type": "keyword"},
			  "Address":  { "type": "text"}, 
			  "AddressSuit": {"type": "text"},
			  "State": {"type": "keyword"},
			  "City":   { "type": "keyword"},
			  "Zip": {"type": "keyword"},
			  "Phone":  { "type": "keyword"}, 
			  "Specialty":   { "type": "keyword"},
			  "SubSpecialty": {"type": "text"},
			  "YearsOfExperience":  { "type": "keyword"},
			  "IsOnlineScheduleEnable": { "type": "boolean"},
              "IsInClinicBookEnable": { "type": "boolean"},
  			  "IsVirtualBookEnable": { "type": "boolean"},
 			  "NextAvailableDateInClinic": {"type": "date"},
              "NextAvailableDateVirtual": {"type": "date"},
			  "Location": {
				"type": "geo_point"
			  }
			}
		}
}
`
)

func SetupElasticSearchEngine()  {
	indexList := [...]elasticIndex {
		{Name: DoctorIndexName, Mapping: DoctorIndexMapping},
	}

	engine := GetElasticSearchEngine()
	for _, index := range indexList {
		createIndexMappingsIfNotExisting(engine, index.Name, index.Mapping)
	}
}

func createIndexMappingsIfNotExisting(client *elastic.Client, name string, mapping string)  {
	exist, err := client.IndexExists(name).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if exist == false {
		index, err := client.CreateIndex(name).Body(mapping).Do(context.TODO())
		if err != nil {
			log.Fatalf("Create Index(%s) fatal.", name)
		}

		if index == nil {
			log.Fatalf("Create Index(%s) fatal, should not nil.", name)
		}
	}
}
