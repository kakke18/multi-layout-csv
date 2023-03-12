package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	recordTypeSummary = "summary"
	recordTypeCountry = "country"
)

type Summary struct {
	RecordType string
	Summary    string
}

type Country struct {
	RecordType string
	Name       string
	ISOCode    string
	Population int
}

type singleCSVReader struct {
	record []string
}

var (
	_ gocsv.CSVReader = singleCSVReader{}
)

func (r singleCSVReader) Read() ([]string, error) {
	return r.record, nil
}

func (r singleCSVReader) ReadAll() ([][]string, error) {
	return [][]string{r.record}, nil
}

func main() {
	s := `summary,3件
country,アメリカ合衆国,US/USA,310232863
country,日本,JP/JPN,127288000
country,中国,CH/CHN,1330044000`

	r := csv.NewReader(strings.NewReader(s)) // 標準パッケージでまず読み込む
	r.FieldsPerRecord = -1                   // レコード数が異なるのでチェックを無効化する
	all, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range all {
		switch record[0] {
		case recordTypeSummary:
			var summaries []Summary
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record}, &summaries); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("summary: %+v\n", summaries[0])
		case recordTypeCountry:
			var countries []Country
			if err := gocsv.UnmarshalCSVWithoutHeaders(singleCSVReader{record}, &countries); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("country: %+v\n", countries[0])
		default:
			log.Fatal("invalid record")
		}
	}
}
