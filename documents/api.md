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
|keyword |yes  |string |   |
|isInClinicEnable |no  |bool |  default true  | 
|isVirtualEnable     |no  |bool |    |
|appointmentType | yes | int | all = 0 inClinic=1 virtual = 2
|nextAvailableDate |no |date | UTC format
|city |no |string |
|specialty |no | string |  
|lat | yes | float | 
|lon |yes | float |
|gender |no | bool|
|page |yes | int| default index from 1
|pageSize |yes | int | default size from 50


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





