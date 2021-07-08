package main

import (
	"OpenSchedule/src/database"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

var elasticSearchEngine *elastic.Client

func main()  {
	database.SetupElasticSearchEngine()
	elasticSearchEngine = database.GetElasticSearchEngine()
	fuzzyQuery("FullName", "kim")
}

func fuzzyQuery(field string, value string)  {
	q := elastic.NewFuzzyQuery(field, value).Boost(1.5).Fuzziness(2).PrefixLength(0).MaxExpansions(100)

	result, err := elasticSearchEngine.Search().Index(database.DoctorIndexName).
		Size(30).
		From(0).
		Query(q).Pretty(true).
		Do(context.Background())

	for _, hit := range result.Hits.Hits {
		var postType struct{FullName string}
		err = json.Unmarshal(hit.Source, &postType)

		if err != nil {
			break
		}

	}

}
