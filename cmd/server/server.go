package main

import (
	"comments-service/pkg/api"
	"comments-service/pkg/storage/postgres"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	PostgresDb string `json:"postgres_db"`
	HostURL    string `json:"host_url"`
}

func main() {
	confFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var conf config
	err = json.Unmarshal(confFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(conf.PostgresDb)
	if err != nil {
		log.Fatal(err)
	}

	apiDb := api.New(db)

	log.Println("[*] HTTP server is started on ", conf.HostURL)
	err = http.ListenAndServe(conf.HostURL, apiDb.Router())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("[*] HTTP server has been stopped. Reason: got sigterm")
	}
}
