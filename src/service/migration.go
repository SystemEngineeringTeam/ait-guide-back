package service

import (
	"context"
	"encoding/csv"
	"fmt"
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

func Migration() {
	locations := model.GetAllPoints()
	if locations == nil {
		records, _ := readCsv("../data/point.csv")
		for i, v := range records {
			if i == 0 {
				continue
			}
			fmt.Println(i, "/", len(records)-1)
			model.InsertPoint(v[0], v[1], v[2])
		}
		CreateGraph(context.Background(), "Point", "Distance")
	}
}
