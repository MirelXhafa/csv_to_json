package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateJSON(filename string, record interface{}) {

	js, err := json.MarshalIndent(record, "", "")

	if err != nil {
		log.Fatal(err)
	}

	_ = ioutil.WriteFile(filename+".json", js, 0644)
}

func GetNote() string {
	response, err := http.Get("http://numbersapi.com/random/math")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)

}
