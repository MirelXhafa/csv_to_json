package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/MirelXhafa/csv_to_ndjson/utils"
)

type TURL struct {
	Scheme string
	Host   string
	Path   string
	Opaque string
}

type TRECORDS struct {
	TS        int64  `json:"ts"`
	Source_ip string `json:"source_ip"`
	Url       TURL   `json:"url"`
	Size      string `json:"size"`
	Note      string `json:"note"`
}

func getCsvData(datas [][]string, records chan []TRECORDS, output_file_name string) {

	record := []TRECORDS{}

	for _, data := range datas {

		for _, dt := range data {
			// split data from csv string
			splitedData := strings.Split(dt, ",")

			// Parse date string with format ISO 8601
			parsedDate, err := time.Parse("2006-01-02T15:04:05Z0700", splitedData[0])

			if err != nil {
				log.Fatal(err)
			}

			// parse url
			url := utils.ParseURL(splitedData[2])

			record = append(record, TRECORDS{parsedDate.Unix(),
				splitedData[1],
				TURL{
					url.Scheme,
					url.Host,
					url.Path,
					url.Opaque,
				},
				splitedData[3],
				utils.GetNote(),
			})
		}

	}
	utils.CreateJSON(output_file_name, record)

	records <- record

	defer close(records)
}

func CSV_TO_JSON(csvFile string, output_file_name string) {
	records := make(chan []TRECORDS)
	file, err := os.Open("./items/example.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parser := csv.NewReader(file)
	data, err := parser.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	go getCsvData(data, records, output_file_name)

	<-records
}
