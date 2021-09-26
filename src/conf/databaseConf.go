package conf

type DatabaseConf struct {
	Host string
	Port int
	User string
	Password string
	DatabaseName string
}

var MySQLConf = DatabaseConf{
	Host: "42.192.92.99",
	Port: 3306,
	User: "gust",
	Password: "Yj202!0701",
	DatabaseName: "openSchedule",
}

var LocalMySqlConf = DatabaseConf{
	Host:"127.0.0.1",
	Port: 3306,
	User:"root",
	Password: "123456",
	DatabaseName:"drfinder",
}

var ElasticSearchConf = DatabaseConf{
	Host: "http://42.192.92.99",
	Port: 9200,
	User: "",
	Password: "",
	DatabaseName: "",
}
