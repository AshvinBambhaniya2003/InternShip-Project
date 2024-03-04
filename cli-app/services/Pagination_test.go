package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID   int
	Name string
}

func TestPaginate(t *testing.T) {
	// Define some test data
	records := []TestStruct{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}

	t.Run("ValidPaginationASC", func(t *testing.T) {
		expectedResult := []TestStruct{{1, "Alice"}, {2, "Bob"}}
		paginationResult, err := Paginate(records, 0, 2, "ID", "ASC")
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, paginationResult)
	})

	t.Run("ValidPaginationDSC", func(t *testing.T) {
		expectedResult := []TestStruct{{2, "Bob"}, {1, "Alice"}}
		paginationResult, err := Paginate(records, 1, 2, "ID", "DSC")
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, paginationResult)
	})

	t.Run("InvalidOrder", func(t *testing.T) {
		_, err := Paginate(records, 0, 2, "ID", "INVALID")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid order specified. Must be 'ASC'|'DSC'")
	})

	t.Run("EmptyRecords", func(t *testing.T) {
		_, err := Paginate([]TestStruct{}, 0, 2, "ID", "ASC")
		assert.Error(t, err)
		assert.EqualError(t, err, "no any records for given field")
	})
}
