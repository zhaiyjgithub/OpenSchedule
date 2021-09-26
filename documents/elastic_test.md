`GET _search
 {
   "query": {
     "match_all": {}
   }
 }
 
 POST /test/user
 {
   "myTime": "2021-07-60 21:05:00"
 }
 
 GET /test/user/_search
 
 PUT /doctor
 {
 		"mappings":{
 			"properties": {
 			  "Npi":    { "type": "long" },  
 			  "FirstName":  { "type": "keyword"}, 
 			  "LastName":   { "type": "keyword"},
 			  "MiddleName": { "type": "keyword"},
 			  "FullName": {"type": "text"},
 			  "NamePrefix":  { "type": "keyword"}, 
 			  "JobTitle":   { "type": "keyword"},
 			  "Gender": {"type": "keyword"},
 			  "Address":  { "type": "text"}, 
 			  "AddressSuit": {"type": "text"},
 			  "State": {"type": "keyword"},
 			  "City":   { "type": "keyword"},
 			  "Zip": {"type": "keyword"},
 			  "Phone":  { "type": "keyword"}, 
 			  "Specialty":   { "type": "keyword"},
 			  "SubSpecialty": {"type": "text"},
 			  "YearsOfExperience":  { "type": "keyword"},
 			  "IsOnlineScheduleEnable": { "type": "boolean"},
               "IsInClinicBookEnable": { "type": "boolean"},
   			  "IsVirtualBookEnable": { "type": "boolean"},
  			  "NextAvailableDateInClinic": {"type": "date"},
               "NextAvailableDateVirtual": {"type": "date"},
 			  "Location": {
 				"type": "geo_point"
 			  }
 			}
 		}
 }
 
 GET /doctor/_mapping
 
 DELETE /doctor
 
 PUT /doctor/_doc/1
 {
   "npi": 1073516027,  
   "fistName": "Chang", 
   "lastName":   "John",
   "fullName": "John C. Chang",
   "namePrefix": "Dr. ", 
   "jobTitle": "MD",
   "gender": "M",
   "address":  "169 N Middletown Rd, Pearl River, NY, 10965", 
   "state": "NY",
   "city":   "Pearl River",
   "zip": "10965",
   "phone":  "(845) 735-5666", 
   "specialty": "Ophthalmology",
   "subSpecialty": "General Ophthalmology, Cataract Related",
   "yearsOfExperience":  "11, 20", 
   "location": {"lat": 41.065736, "lon": -74.012544}
 }
 
 GET /doctor/_search
 {
   "query": {
     "bool": {
       "must": {
         "match_all": {}
       },
       "filter": {
         "geo_distance": {
           "distance": "2000km",
           "Location": {
             "lat": 40,
             "lon": -70
           }
         }
       }
     }
   }
 }
 
 GET /doctor/_count
 
 GET /doctor/_mapping
 
 POST /doctor/_search
 {
     "query": {
         "multi_match" : {
             "query" : "gold",
             "fields": ["FullName"],
             "fuzziness": "AUTO"
         }
     }
 }
 
 # doctor name fuzzy search
 GET /doctor/_search
 {
   "query": {
     "bool": {
       "must": [
         {"match" : {
            "FullName": {
              "query": "sliver",
              "fuzziness": "AUTO"
            }
         }}
       ],
       "filter": [
         {"term": {
           "Gender": "F"
         }}
       ]
     }
   }
 }
 
 #multi term prices search
 GET /doctor/_search
 {
   "query": {
     "bool": {
       "must": [
         {"term": {
           "City": "Babylon"
         }},
         {"term": {
           "Zip": "11702"
         }},
         {"term": {
           "Phone": "(631) 321-2100"
         }},
         {"range": {
           "NextAvailableDateInClinic": {
             "gte": "2021-07-07T14:36:41Z"
           }
         }}
       ],
       "filter": [
         {"term": {
           "Gender": "F"
         }}
       ]
     }
   }
 }
 
 #multi term and time range prices search
 GET /doctor/_search
 {
   "query": {
     "bool": {
       "must": [
         {"term": {
           "City": "Babylon"
         }},
         {"term": {
           "Zip": "11702"
         }},
         {"term": {
           "Phone": "(631) 321-2100"
         }}
       ],
       "filter": [
         {"term": {
           "Gender": "F"
         }}
       ]
     }
   }
 }
 
 #multi term geo prices search
 GET /doctor/_search
 {
   "query": {
     "bool": {
       "must": [
         {"term": {
           "City": "Babylon"
         }},
         {"term": {
           "Zip": "11702"
         }},
         {"range": {
           "NextAvailableDateInClinic": {
             "gte": "2021-07-05T14:36:41Z"
           }
         }}
       ],
       "filter": [
         {"term": {
           "Gender": "F"
         }},
          {"geo_distance": {
           "distance": "200km",
           "Location": {
             "lat": 40.747898,
             "lon": -73.324025
           }
         }}
       ]
     }
   },
   "sort": [
     {
       "_geo_distance": {
         "Location": { 
           "lat":  40.747898,
           "lon": -73.324025
         },
         "order":         "asc",
         "unit":          "km", 
         "distance_type": "plane" 
       }
     }
   ],
   "from": 10,
   "size": 20
 }`