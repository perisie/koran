package utils

import (
	"strconv"

	"github.com/arikama/koran-backend/beans"
)

func ReadSurahsInfo() ([]*beans.SurahInfo, error) {
	records, err := ReadQuranCsv("surahs")
	if err != nil {
		return nil, err
	}
	result := []*beans.SurahInfo{}
	for _, record := range records {
		surahId, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		verses, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, err
		}
		juzStart, err := strconv.Atoi(record[6])
		if err != nil {
			return nil, err
		}
		juzEnd, err := strconv.Atoi(record[7])
		if err != nil {
			return nil, err
		}
		surahInfo := beans.SurahInfo{
			SurahId:  surahId,
			Title:    record[1],
			Arabic:   record[2],
			English:  record[3],
			Verses:   verses,
			City:     record[5],
			JuzStart: juzStart,
			JuzEnd:   juzEnd,
		}
		result = append(result, &surahInfo)
	}
	return result, nil
}
