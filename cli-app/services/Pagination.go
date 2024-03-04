package services

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

func Paginate[T any](records []T, skip, limit int, orderBy string, order string) ([]T, error) {

	if len(records) == 0 {
		log.Fatal("no any records for given field")
	}
	if skip < 0 {
		skip = 0
	}
	if limit < -1 {
		limit = -1
	}

	if orderBy != "" {
		orderLower := strings.ToLower(order)
		if orderLower != "asc" && orderLower != "dsc" {
			return nil, errors.New("invalid order specified. Must be 'ASC'|'DSC'")
		}

		sort.SliceStable(records, func(i, j int) bool {
			if orderLower == "asc" {
				return getField(records[i], orderBy) < getField(records[j], orderBy)
			} else if orderLower == "dsc" {
				return getField(records[i], orderBy) > getField(records[j], orderBy)
			}
			return false
		})
	}

	end := len(records)

	if skip < 0 || end < 0 || skip >= len(records) || end > len(records) {
		return nil, errors.New("invalid pagination indices")
	}

	if limit != -1 && skip+limit < end {
		end = skip + limit
	}

	return records[skip:end], nil
}

func SelectColumn[T any](records []T, selects string) []map[string]interface{} {

	selectedColumns := strings.Split(selects, ",")
	var selectedRecords []map[string]interface{}
	for _, record := range records {
		selectedRecord := make(map[string]interface{})
		for _, selectField := range selectedColumns {
			selectedRecord[selectField] = getField(record, selectField)
		}
		selectedRecords = append(selectedRecords, selectedRecord)
	}

	return selectedRecords
}

func getField(record interface{}, field string) string {
	r := reflect.ValueOf(record)
	f := reflect.Indirect(r).FieldByName(field)
	if f.IsValid() {
		switch f.Kind() {
		case reflect.String:
			return f.Interface().(string)
		case reflect.Int:
			return fmt.Sprintf("%d", f.Interface())
		default:
			return fmt.Sprintf("%v", f.Interface())
		}
	}
	return ""
}
