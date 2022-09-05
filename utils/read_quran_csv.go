package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ReadQuranCsv(filename string) ([][]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%v/../qurancsv/%v.csv", wd, filename)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// https://stackoverflow.com/questions/21371673/reading-files-with-a-bom-in-go
	if len(records) > 0 && len(records[0]) > 0 {
		records[0][0] = strings.Trim(records[0][0], "\ufeff")
	}

	return records, nil
}
