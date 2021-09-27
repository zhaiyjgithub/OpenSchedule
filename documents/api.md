Open API

    
##### 简要描述

- 模糊搜索

##### 请求URL
- ` http://xx.com/api/search/fuzzySearch `
  
##### 请求方式
- POST 

##### 参数

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Keyword |yes  |string |   |
|IsInClinicEnable |no  |bool |  default true  | 
|IsVirtualEnable     |no  |bool |    |
|AppointmentType | yes | int | all = 0 inClinic=1 virtual = 2
|NextAvailableDate |no |date | UTC format
|City |no |string |
|Specialty |no | string |  
|Lat | yes | float | 
|Lon |yes | float |
|Gender |no | bool|
|Page |yes | int| default index from 1
|PageSize |yes | int | default size from 50
|Distance |no | int | default 200
|SortType |no | int | byDefault = 0, ByDistance = 1


##### Response sample

``` 
  {
    "code": 0,
    "data": [
          "Npi" : 1992729248,
          "LastName" : "Uyar",
          "FirstName" : "Conchita",
          "MiddleName" : "",
          "FullName" : "Conchita A. Uyar",
          "NamePrefix" : "Dr. ",
          "JobTitle" : "MD",
          "Gender" : "F",
          "Address" : "2 Burnage Ln, Babylon, NY, 11702",
          "City" : "Babylon",
          "State" : "NY",
          "Zip" : "11702",
          "Phone" : "(631) 587-7641",
          "Specialty" : "Obstetrics & Gynecology",
          "SubSpecialty" : "General Obstetrics & Gynecology",
          "AddressSuit" : "",
          "YearsOfExperience" : "21, -1",
          "Location" : {
            "lat" : 40.713609,
            "lon" : -73.321625
          },
          "IsOnlineScheduleEnable" : true,
          "IsInClinicBookEnable" : true,
          "IsVirtualBookEnable" : true,
          "NextAvailableDateInClinic" : "2021-07-07T15:57:48Z",
          "NextAvailableDateVirtual" : "2021-07-07T15:57:48Z",
          "Distance": 3.8181313249362967
    ]
  }
```

##### Response

|Name|Type|Description|
|:-----  |:-----|-----                           |
|Distance |float   | return 0 if lat or lon  is not specified in params|

##### 备注 

- 更多返回错误代码请看首页的错误代码描述


-------

##### 简要描述

- get the settings of schedule

##### 请求URL
- ` http://xx.com/api/schedule/getSettings `
  
##### 请求方式
- POST 

##### 参数

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|npi |yes |int64 | 

##### Response sample

``` 
  {
    "code": 0,
    "data": [
          "Npi" : 1992729248,
          "LastName" : "Uyar",
          "FirstName" : "Conchita",
          "MiddleName" : "",
          "FullName" : "Conchita A. Uyar",
          "NamePrefix" : "Dr. ",
          "JobTitle" : "MD",
          "Gender" : "F",
          "Address" : "2 Burnage Ln, Babylon, NY, 11702",
          "City" : "Babylon",
          "State" : "NY",
          "Zip" : "11702",
          "Phone" : "(631) 587-7641",
          "Specialty" : "Obstetrics & Gynecology",
          "SubSpecialty" : "General Obstetrics & Gynecology",
          "AddressSuit" : "",
          "YearsOfExperience" : "21, -1",
          "Location" : {
            "lat" : 40.713609,
            "lon" : -73.321625
          },
          "IsOnlineScheduleEnable" : true,
          "IsInClinicBookEnable" : true,
          "IsVirtualBookEnable" : true,
          "NextAvailableDateInClinic" : "2021-07-07T15:57:48Z",
          "NextAvailableDateVirtual" : "2021-07-07T15:57:48Z",
          "Distance": 3.8181313249362967
    ]
  }
```

##### Response

|Name|Type|Description|
|:-----  |:-----|-----                           |
|Distance |float   | return 0 if lat or lon  is not specified in params|

##### 备注 

- 更多返回错误代码请看首页的错误代码描述


-------


##### 简要描述

- Set the schedule settings.

##### 请求URL
- ` http://xx.com/api/schedule/setScheduleSettings `

##### 请求方式
- POST

##### 参数

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|npi |yes |int64 |
|duration_per_slot |yes |int64 |
|number_per_slot |yes |int64 |
|monday_am_is_enable |yes |bool |
|monday_am_start_time |yes |dateTime |
|monday_am_end_time |yes |dateTime |
|monday_am_appointment_type |yes |int |
|monday_pm_is_enable |yes |bool |
|monday_pm_start_time |yes |dateTime |
|monday_pm_end_time |yes |dateTime |
|monday_pm_appointment_type |yes |int |
|tuesday_am_is_enable |yes |bool |
|tuesday_am_start_time |yes |dateTime |
|tuesday_am_end_time |yes |dateTime |
|tuesday_am_appointment_type |yes |int |
|tuesday_pm_is_enable |yes |bool |
|tuesday_pm_start_time |yes |dateTime |
|tuesday_pm_end_time |yes |dateTime |
|tuesday_pm_appointment_type |yes |int |
|wednesday_am_is_enable |yes |bool |
|wednesday_am_start_time |yes |dateTime |
|wednesday_am_end_time |yes |dateTime |
|wednesday_am_appointment_type |yes |int |
|wednesday_pm_is_enable |yes |bool |
|wednesday_pm_start_time |yes |dateTime |
|wednesday_pm_end_time |yes |dateTime |
|wednesday_pm_appointment_type |yes |int |
|thursday_am_is_enable |yes |bool |
|thursday_am_start_time |yes |dateTime |
|thursday_am_end_time |yes |dateTime |
|thursday_am_appointment_type |yes |int |
|thursday_pm_is_enable |yes |bool |
|thursday_pm_start_time |yes |dateTime |
|thursday_pm_end_time |yes |dateTime |
|thursday_pm_appointment_type |yes |int |
|friday_am_is_enable |yes |bool |
|friday_am_start_time |yes |dateTime |
|friday_am_end_time |yes |dateTime |
|friday_am_appointment_type |yes |int |
|friday_pm_is_enable |yes |bool |
|friday_pm_start_time |yes |dateTime |
|friday_pm_end_time |yes |dateTime |
|friday_pm_appointment_type |yes |int |
|saturday_am_is_enable |yes |bool |
|saturday_am_start_time |yes |dateTime |
|saturday_am_end_time |yes |dateTime |
|saturday_am_appointment_type |yes |int |
|saturday_pm_is_enable |yes |bool |
|saturday_pm_start_time |yes |dateTime |
|saturday_pm_end_time |yes |dateTime |
|saturday_pm_appointment_type |yes |int |
|sunday_am_is_enable |yes |bool |
|sunday_am_start_time |yes |dateTime |
|sunday_am_end_time |yes |dateTime |
|sunday_am_appointment_type |yes |int |
|sunday_pm_is_enable |yes |bool |
|sunday_pm_start_time |yes |dateTime |
|sunday_pm_end_time |yes |dateTime |
|sunday_pm_appointment_type |yes |int |

##### Response sample

``` 
  {
    "code": 0,
    "data": true,
    "msg": "success"
  }
```

##### Response

|Name|Type|Description|
|:-----  |:-----|-----                           |
|Distance |float   | return 0 if lat or lon  is not specified in params|

##### 备注

- 更多返回错误代码请看首页的错误代码描述


-------




