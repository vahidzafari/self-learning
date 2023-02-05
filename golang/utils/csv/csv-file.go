package utils

import (
	"encoding/csv"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func ReadCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// CSV file read all at once
	// lines data type is [][]string
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func SaveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}
