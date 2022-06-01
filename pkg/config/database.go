package config

import "fmt"

const DB_NAME = "basic"
const DB_USER = "admin"
const DB_PASSWORD = "12345678"
const DB_HOST = "127.0.0.1"
const DB_PORT = "5432"
const DB_SSL_MODE = "disable"

func DBConfiguration() string {
	dbDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE)

	return dbDsn
}
