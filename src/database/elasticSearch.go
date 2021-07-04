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
			  "npi":    { "type": "long" },  
			  "fistName":  { "type": "keyword"}, 
			  "lastName":   { "type": "keyword"},
			  "fullName": {"type": "keyword"},
			  "namePrefix":  { "type": "keyword"}, 
			  "jobTitle":   { "type": "keyword"},
			  "gender": {"type": "keyword"},
			  "address":  { "type": "text"}, 
			  "state": {"type": "keyword"},
			  "city":   { "type": "keyword"},
			  "zip": {"type": "keyword"},
			  "phone":  { "type": "keyword"}, 
			  "specialty":   { "type": "keyword"},
			  "subSpecialty": {"type": "keyword"},
			  "yearsOfExperience":  { "type": "keyword"}, 
			  "pin": {
				"properties": {
				  "location": {
					"type": "geo_point"
				  }
				}
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
