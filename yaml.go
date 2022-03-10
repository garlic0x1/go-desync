package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Template struct {
	Oracle  []string
	Request string
}

func parseRequest(file string) string {
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data["Request"]
}

func parseMatcher(file string) string {
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data["Match"]
}
