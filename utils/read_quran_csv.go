package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadQuranCsv(filename string) ([][]string, error) {
	wd, err := os.Getwd()
	before, _, ok := strings.Cut(wd, "koran-backend")
	if !ok {
		return nil, errors.New("invalid working directory")
	}
	basePath := before + "koran-backend/qurancsv"
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%v/%v.csv", basePath, filename)
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
