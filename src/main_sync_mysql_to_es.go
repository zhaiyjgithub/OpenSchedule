package main

import (
	"OpenSchedule/src/database"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

var mysqlEngine *gorm.DB
var bulkService *elastic.BulkService

type Doctor struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;"`
	//[ 1] npi                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Npi int64 `gorm:"column:npi;type:int;"`
	//[ 2] last_name                                      varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LastName string `gorm:"column:last_name;type:varchar;size:50;"`
	//[ 3] first_name                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	FirstName string `gorm:"column:first_name;type:varchar;size:50;"`
	//[ 4] middle_name                                    varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	MiddleName string `gorm:"column:middle_name;type:varchar;size:20;"`
	//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	//[ 6] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
	//[ 7] full_name                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullName string `gorm:"column:full_name;type:varchar;size:255;"`
	//[ 8] name_prefix                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	NamePrefix string `gorm:"column:name_prefix;type:varchar;size:10;"`
	//[ 9] credential                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	JobTitle string `gorm:"column:credential;type:varchar;size:50;"`
	//[10] gender                                         varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	Gender string `gorm:"column:gender;type:varchar;size:2;"`
	//[11] address                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Address string `gorm:"column:address;type:text;size:65535;"`
	//[12] city                                           varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	City string `gorm:"column:city;type:varchar;size:20;"`
	//[13] state                                          varchar(5)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 5       default: []
	State string `gorm:"column:state;type:varchar;size:5;"`
	//[14] zip                                            varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Zip string `gorm:"column:zip;type:varchar;size:20;"`
	//[15] phone                                          varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Phone string `gorm:"column:phone;type:varchar;size:20;"`
	//[16] specialty                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Specialty string `gorm:"column:specialty;type:varchar;size:255;"`
	//[17] sub_specialty                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SubSpecialty string `gorm:"column:sub_specialty;type:varchar;size:255;"`
	//[21] address_suit                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AddressSuit string `gorm:"column:address_suit;type:varchar;size:20;"`
	//[22] lang                                           varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Lang string `gorm:"column:lang;type:varchar;size:50;"`
	//[23] year_of_experience                             varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	YearOfExperience string `gorm:"column:year_of_experience;type:varchar;size:10;"`
	Lat float64
	Lng float64
}

type ESDoctor struct {
	//[ 1] npi                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Npi int64 `gorm:"column:npi;type:int;"`
	//[ 2] last_name                                      varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LastName string `gorm:"column:last_name;type:varchar;size:50;"`
	//[ 3] first_name                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	FirstName string `gorm:"column:first_name;type:varchar;size:50;"`
	//[ 4] middle_name                                    varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	MiddleName string `gorm:"column:middle_name;type:varchar;size:20;"`
	//[ 7] full_name                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullName string `gorm:"column:full_name;type:varchar;size:255;"`
	//[ 8] name_prefix                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	NamePrefix string `gorm:"column:name_prefix;type:varchar;size:10;"`
	//[ 9] credential                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	JobTitle string `gorm:"column:credential;type:varchar;size:50;"`
	//[10] gender                                         varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	Gender string `gorm:"column:gender;type:varchar;size:2;"`
	//[11] address                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Address string `gorm:"column:address;type:text;size:65535;"`
	//[12] city                                           varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	City string `gorm:"column:city;type:varchar;size:20;"`
	//[13] state                                          varchar(5)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 5       default: []
	State string `gorm:"column:state;type:varchar;size:5;"`
	//[14] zip                                            varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Zip string `gorm:"column:zip;type:varchar;size:20;"`
	//[15] phone                                          varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Phone string `gorm:"column:phone;type:varchar;size:20;"`
	//[16] specialty                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Specialty string `gorm:"column:specialty;type:varchar;size:255;"`
	//[17] sub_specialty                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SubSpecialty string `gorm:"column:sub_specialty;type:varchar;size:255;"`
	//[21] address_suit                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AddressSuit string `gorm:"column:address_suit;type:varchar;size:20;"`
	//[23] year_of_experience                             varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	YearsOfExperience string `gorm:"column:years_of_experience;type:varchar;size:10;"`
	Location elastic.GeoPoint
	IsOnlineScheduleEnable bool
	IsInClinicBookEnable bool
	IsVirtualBookEnable bool
	NextAvailableDateInClinic string
	NextAvailableDateVirtual string
}

func main()  {
	database.SetupElasticSearchEngine()
	mysqlEngine = database.GetLocalMySqlEngine()
	bulkService = database.GetElasticSearchEngine().Bulk()

	const pageSize = 1000
	page := 1
	for {
		doctors := getDoctor(page, pageSize)
		bulkDoctors(doctors)
		if len(doctors) < pageSize {
			fmt.Println("least page: ", page)
			break
		}
		fmt.Printf("page %d is finished ...\n", page)
		page = page + 1
		time.Sleep(time.Microsecond*500)
	}
}

func getDoctor(page int, pageSize int) []*Doctor  {
	var doctors []*Doctor
	db:= mysqlEngine.Raw("SELECT doctors.*, doctors.credential as job_title, geos.lat, geos.lng from doctors left JOIN geos on geos.npi = doctors.npi limit ? offset ?", pageSize, (page -1)*pageSize).Scan(&doctors)
	if db.Error != nil {
		log.Fatal(db.Error)
	}

	return doctors
}

func bulkDoctors(doctors []*Doctor)  {
	for _, doctor := range doctors {
		esDoctor := ESDoctor{}
		esDoctor.Npi = doctor.Npi
		esDoctor.FirstName = doctor.FirstName
		esDoctor.LastName = doctor.LastName
		esDoctor.FullName = doctor.FullName
		esDoctor.MiddleName = doctor.MiddleName
		esDoctor.NamePrefix = doctor.NamePrefix
		esDoctor.Gender = doctor.Gender
		esDoctor.JobTitle = doctor.JobTitle
		esDoctor.AddressSuit = doctor.AddressSuit
		esDoctor.Address = doctor.Address
		esDoctor.City = doctor.City
		esDoctor.State = doctor.State
		esDoctor.Zip = doctor.Zip
		esDoctor.Phone = doctor.Phone
		esDoctor.Specialty = doctor.Specialty
		esDoctor.SubSpecialty = doctor.SubSpecialty
		esDoctor.YearsOfExperience = doctor.YearOfExperience
		esDoctor.Location = elastic.GeoPoint{Lat: doctor.Lat, Lon: doctor.Lng}
		esDoctor.IsOnlineScheduleEnable = true
		esDoctor.IsInClinicBookEnable = true
		esDoctor.IsVirtualBookEnable = true

		date := time.Now().UTC().Format(time.RFC3339)
		esDoctor.NextAvailableDateInClinic = date
		esDoctor.NextAvailableDateVirtual = date

		req := elastic.NewBulkCreateRequest().Index(database.DoctorIndexName).Doc(esDoctor)
		bulkService = bulkService.Add(req)

		bulkResponse, err := bulkService.Do(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		if bulkResponse == nil {
			log.Fatal("expected bulkResponse to be != nil; got nil")
		}

		if bulkService.NumberOfActions() != 0 {
			log.Fatalf("expected bulkService.NumberOfActions %d; got %d", 0, bulkService.NumberOfActions())
		}
	}
}