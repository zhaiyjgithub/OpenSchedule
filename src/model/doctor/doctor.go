package doctor

import (
	"time"
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
type Doctor struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;"json:"-"`
	//[ 1] npi                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Npi int64 `gorm:"column:npi;type:int;"`
	//[ 2] last_name                                      varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LastName string `gorm:"column:last_name;type:varchar;size:50;"`
	//[ 3] first_name                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	FirstName string `gorm:"column:first_name;type:varchar;size:50;"`
	//[ 4] middle_name                                    varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	MiddleName string `gorm:"column:middle_name;type:varchar;size:20;"`
	//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"json:"-"`
	//[ 6] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"json:"-"`
	//[ 7] full_name                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FullName string `gorm:"column:full_name;type:varchar;size:255;"`
	//[ 8] name_prefix                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	NamePrefix string `gorm:"column:name_prefix;type:varchar;size:10;"`
	//[ 9] credential                                     varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Credential string `gorm:"column:credential;type:varchar;size:50;"`
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
	//[18] job_title                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	JobTitle string `gorm:"column:job_title;type:varchar;size:100;"`
	//[19] summary                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Summary string `gorm:"column:summary;type:varchar;size:255;"`
	//[20] fax                                            varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Fax string `gorm:"column:fax;type:varchar;size:20;"`
	//[21] address_suit                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AddressSuit string `gorm:"column:address_suit;type:varchar;size:20;"`
	//[22] lang                                           varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Lang string `gorm:"column:lang;type:varchar;size:50;"`
	//[23] year_of_experience                             varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	YearOfExperience string `gorm:"column:year_of_experience;type:varchar;size:10;"`
}

// TableName sets the insert table name for this struct type
func (d *Doctor) TableName() string {
	return "doctors"
}