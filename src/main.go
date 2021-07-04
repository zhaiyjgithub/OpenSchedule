package main
import (
	"OpenSchedule/src/database"
	"fmt"
)


func main()  {
	fmt.Println("hello go")
	database.SetupElasticSearchEngine()
}
