package generateModel

import (
	"database/sql"
	"github.com/google/uuid"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `doctors` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `npi` int(20) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `middle_name` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `name_prefix` varchar(10) DEFAULT NULL,
  `credential` varchar(50) DEFAULT NULL,
  `gender` varchar(2) DEFAULT NULL,
  `address` text,
  `city` varchar(20) DEFAULT NULL,
  `state` varchar(5) DEFAULT NULL,
  `zip` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `specialty` varchar(255) DEFAULT NULL,
  `sub_specialty` varchar(255) DEFAULT NULL,
  `job_title` varchar(100) DEFAULT NULL,
  `summary` varchar(255) DEFAULT NULL,
  `fax` varchar(20) DEFAULT NULL,
  `address_suit` varchar(20) DEFAULT NULL,
  `lang` varchar(50) DEFAULT NULL,
  `year_of_experience` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_last_name` (`last_name`),
  KEY `idx_spectialty` (`specialty`),
  KEY `idx_first_name` (`first_name`)
) ENGINE=InnoDB AUTO_INCREMENT=238106 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 1,    "npi": 55,    "last_name": "qEBIpBquuhbeOylUqcvwEDfqr",    "first_name": "AqTYReZgRSGBgIMTNsSiVSFPX",    "middle_name": "eiQLDSuPjaGwqjGMjyxhSIVby",    "created_at": "2226-06-01T02:13:23.626256822+08:00",    "updated_at": "2271-10-04T01:49:13.390165631+08:00",    "full_name": "xJrnIXXCeUoCufdfmkCVOrngq",    "name_prefix": "HusCkEauPqpxZIvJsmqnmHnKc",    "credential": "YhIKiqNpQCXhhQTaKBxWHNxiU",    "gender": "yTXCmIigEEvyYJuLqlvEhIvaw",    "address": "XLNdLiqTfsZRrrXjPVcLNeCxk",    "city": "iFPcnNenhEXWRvkBIoTSVMEpo",    "state": "ZOllQPLQUAkmbbTnZpdmXaijt",    "zip": "swRDytYxJJDOPEScBATJkwrYw",    "phone": "nKLoYIlQeuXQAFeuGSDnXJmGi",    "specialty": "nQpOKvPBFESUdyvExXlvrQlAc",    "sub_specialty": "fSjxTkotciDqpIBWNksGwXHLY",    "job_title": "OQAfFIqSHaFbuoUBTZbsvAjcX",    "summary": "BQCYJYoGWCOruNvMnKdckobSQ",    "fax": "FpvVYhAfrmveQPnFmdEHHKNAS",    "address_suit": "SUorpPsppAgVpPghPVblGYfsc",    "lang": "HGNLNbyXPMRLFXsptuEXQPwPQ",    "year_of_experience": "bBuPceXVJgNFqPgcujSTHsGZd"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Doctors struct is a row record of the doctors table in the drfinder database
type Doctors struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] npi                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Npi sql.NullInt64 `gorm:"column:npi;type:int;" json:"npi"`
	//[ 2] last_name                                      varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LastName sql.NullString `gorm:"column:last_name;type:varchar;size:50;" json:"last_name"`
	//[ 3] first_name                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	FirstName sql.NullString `gorm:"column:first_name;type:varchar;size:50;" json:"first_name"`
	//[ 4] middle_name                                    varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	MiddleName sql.NullString `gorm:"column:middle_name;type:varchar;size:20;" json:"middle_name"`
	//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 6] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	//[ 7] full_name                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullName sql.NullString `gorm:"column:full_name;type:varchar;size:255;" json:"full_name"`
	//[ 8] name_prefix                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	NamePrefix sql.NullString `gorm:"column:name_prefix;type:varchar;size:10;" json:"name_prefix"`
	//[ 9] credential                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Credential sql.NullString `gorm:"column:credential;type:varchar;size:50;" json:"credential"`
	//[10] gender                                         varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	Gender sql.NullString `gorm:"column:gender;type:varchar;size:2;" json:"gender"`
	//[11] address                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Address sql.NullString `gorm:"column:address;type:text;size:65535;" json:"address"`
	//[12] city                                           varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	City sql.NullString `gorm:"column:city;type:varchar;size:20;" json:"city"`
	//[13] state                                          varchar(5)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 5       default: []
	State sql.NullString `gorm:"column:state;type:varchar;size:5;" json:"state"`
	//[14] zip                                            varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Zip sql.NullString `gorm:"column:zip;type:varchar;size:20;" json:"zip"`
	//[15] phone                                          varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Phone sql.NullString `gorm:"column:phone;type:varchar;size:20;" json:"phone"`
	//[16] specialty                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Specialty sql.NullString `gorm:"column:specialty;type:varchar;size:255;" json:"specialty"`
	//[17] sub_specialty                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SubSpecialty sql.NullString `gorm:"column:sub_specialty;type:varchar;size:255;" json:"sub_specialty"`
	//[18] job_title                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	JobTitle sql.NullString `gorm:"column:job_title;type:varchar;size:100;" json:"job_title"`
	//[19] summary                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Summary sql.NullString `gorm:"column:summary;type:varchar;size:255;" json:"summary"`
	//[20] fax                                            varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Fax sql.NullString `gorm:"column:fax;type:varchar;size:20;" json:"fax"`
	//[21] address_suit                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AddressSuit sql.NullString `gorm:"column:address_suit;type:varchar;size:20;" json:"address_suit"`
	//[22] lang                                           varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Lang sql.NullString `gorm:"column:lang;type:varchar;size:50;" json:"lang"`
	//[23] year_of_experience                             varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	YearOfExperience sql.NullString `gorm:"column:year_of_experience;type:varchar;size:10;" json:"year_of_experience"`
}

var doctorsTableInfo = &TableInfo{
	Name: "doctors",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "npi",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Npi",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "npi",
			ProtobufFieldName:  "npi",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "last_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "LastName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "last_name",
			ProtobufFieldName:  "last_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "first_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "FirstName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "first_name",
			ProtobufFieldName:  "first_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "middle_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "MiddleName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "middle_name",
			ProtobufFieldName:  "middle_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "full_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "FullName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "full_name",
			ProtobufFieldName:  "full_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "name_prefix",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "NamePrefix",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name_prefix",
			ProtobufFieldName:  "name_prefix",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "credential",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Credential",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "credential",
			ProtobufFieldName:  "credential",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "gender",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "Gender",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "address",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Address",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "city",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "City",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "state",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       5,
			GoFieldName:        "State",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "state",
			ProtobufFieldName:  "state",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "zip",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Zip",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "zip",
			ProtobufFieldName:  "zip",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "phone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Phone",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "specialty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Specialty",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "specialty",
			ProtobufFieldName:  "specialty",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "sub_specialty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SubSpecialty",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sub_specialty",
			ProtobufFieldName:  "sub_specialty",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "job_title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "JobTitle",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "job_title",
			ProtobufFieldName:  "job_title",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "summary",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Summary",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "summary",
			ProtobufFieldName:  "summary",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "fax",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Fax",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "fax",
			ProtobufFieldName:  "fax",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "address_suit",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "AddressSuit",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "address_suit",
			ProtobufFieldName:  "address_suit",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "lang",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Lang",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "lang",
			ProtobufFieldName:  "lang",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "year_of_experience",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "YearOfExperience",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "year_of_experience",
			ProtobufFieldName:  "year_of_experience",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Doctors) TableName() string {
	return "doctors"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Doctors) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Doctors) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Doctors) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Doctors) TableInfo() *TableInfo {
	return doctorsTableInfo
}
