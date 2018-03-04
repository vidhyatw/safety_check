package config

import "os"

type Config struct {
}

func GetEnvironment() string {
	return os.Getenv("env")
}

func GetDatabaseHost() string {
	return os.Getenv("dbhost")
}

func GetCollection() string {
	return os.Getenv("collection")
}
