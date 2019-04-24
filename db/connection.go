package db

import (
	"github.com/lmorais/supernova/config"
	"log"
)

func Connect(uri string, dbname string) (*Session, error) {

	if uri == "" && dbname == "" {
		session, err := NewSession(config.Fetch().MongoURI, config.Fetch().MongoDatabase)
		if err != nil {
			log.Fatalf("Failed to connect to mongodb: %s", err)
		}

		return session, nil
	}

	session, err := NewSession(uri, dbname)

	if err != nil {
		log.Fatalf("Failed to connect to mongodb: %s", err)
	}

	return session, nil

}
