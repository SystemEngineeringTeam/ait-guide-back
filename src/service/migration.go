package service

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/SystemEngineeringTeam/ait-guide-back/model"
)

func readCsv(filename string) ([][]string, error) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		defer file.Close()
		reader := csv.NewReader(file)
		if records, err := reader.ReadAll(); err != nil {
			log.Fatal(err)
			return nil, err
		} else {
			return records, nil
		}
	}
}

func Migration(yymmdd string) {
	records, _ := readCsv("data/points" + yymmdd + ".csv")
	locations := model.GetAllPoints()
	if locations == nil {
		for _, v := range records {
			model.InsertPoint(v[0], v[1], v[2])
		}
	}
}
