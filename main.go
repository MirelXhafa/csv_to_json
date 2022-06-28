package main

import (
	"github.com/MirelXhafa/csv_to_ndjson/csv"
)

func main() {

	csv.CSV_TO_JSON("./items/example.csv", "network_data")

}
