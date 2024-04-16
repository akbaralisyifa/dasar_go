package database

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabse() string {
	return connection
}