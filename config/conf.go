package config

import (
	"os"
)

type Conf struct {
	MongoURI      string
	MongoDatabase string
	APIPort       string
}

func Fetch() Conf {
	conf := Conf{
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
		APIPort:       ":" + os.Getenv("API_PORT"),
	}

	return conf
}
