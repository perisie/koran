package quran

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_quran_csv(file_path string) ([][]string, error) {
	file, err := os.Open(file_path)
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

func surah_verse_key(surah_id, verse_id int) string {
	return fmt.Sprintf("%v:%v", surah_id, verse_id)
}

func read_surah_info(file_path string) ([]*Surah_info, error) {
	records, err := read_quran_csv(file_path)
	if err != nil {
		return nil, err
	}
	var result []*Surah_info
	for _, record := range records {
		surah_id, err_record := strconv.Atoi(record[0])
		if err_record != nil {
			return nil, err_record
		}
		verses, err_record := strconv.Atoi(record[4])
		if err_record != nil {
			return nil, err_record
		}
		juz_start, err_record := strconv.Atoi(record[6])
		if err_record != nil {
			return nil, err_record
		}
		juz_end, err_record := strconv.Atoi(record[7])
		if err_record != nil {
			return nil, err_record
		}
		surahInfo := Surah_info{
			Surah_id:  surah_id,
			Title:     record[1],
			Arabic:    record[2],
			English:   record[3],
			Verses:    verses,
			City:      record[5],
			Juz_start: juz_start,
			Juz_end:   juz_end,
		}
		result = append(result, &surahInfo)
	}
	return result, nil
}
