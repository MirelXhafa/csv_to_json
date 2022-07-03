package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	ndjson "github.com/scizorman/go-ndjson"
)

func CreateJSON(filename string, record interface{}) {

	r, err := ndjson.Marshal(record)

	if err != nil {
		fmt.Print(err)
	}

	_ = ioutil.WriteFile(filename+".ndjson", r, 0644)
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
