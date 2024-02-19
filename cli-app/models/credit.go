package models

import (
	"encoding/csv"
	"os"
	"strings"
)


type Credit struct {
	PersonID  int    
	TitleID   string 
	Name      string 
	Character string 
	Role      string 
}

func ReadCredits(filename string) ([]Credit, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var credits []Credit
	first := true
	for _, record := range records {

		if first {
			first = false
			continue
		}
		credit := Credit{
			PersonID: parseInt(record[0]),
			TitleID:  record[1],
			Name:     record[2],
			Character: strings.Trim(record[3], `"`),
			Role:     record[4],
		}
		
		credits = append(credits, credit)
		
		
	}

	return credits, nil
}