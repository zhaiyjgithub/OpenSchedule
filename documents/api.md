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





