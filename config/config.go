package config

import "os"

func SecretKey() string {
	return os.Getenv("SECRET_JWT")
}
