package database

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	URL      string
	NAME     string
	PASSWORD string
}

func InitDatabase(conf Config) *elasticsearch.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf.URL = os.Getenv("URL")
	conf.NAME = os.Getenv("NAME")
	conf.PASSWORD = os.Getenv("PASSWORD")

	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.URL,
		},
		Username: conf.NAME,
		Password: conf.PASSWORD,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return es
}
