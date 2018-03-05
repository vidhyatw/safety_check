package config

import "os"

type Config struct {
}

func GetEnvironment() string {
	return os.Getenv("env")
}

func GetDatabaseURI() string {
	return os.Getenv("MONGODB_URI")
}
