package main

import (
	"OpenSchedule/src/database"
	"fmt"
)


func main()  {
	fmt.Println("hello go")
	database.SetupElasticSearchEngine()

	//fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
