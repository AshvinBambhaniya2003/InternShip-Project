package services

import (
	"netflix/models"
	"testing"
)

func TestListTitleWithPersonCount(t *testing.T) {
	// Sample data
	titles := []models.Title{
		{ID: "1", Title: "Title 1"},
		{ID: "2", Title: "Title 2"},
		{ID: "3", Title: "Another"},
	}
	credits := []models.Credit{
		{TitleID: "1", PersonID: 101},
		{TitleID: "1", PersonID: 102},
		{TitleID: "2", PersonID: 103},
		{TitleID: "3", PersonID: 104},
	}

	// Test cases
	testCases := []struct {
		name        string
		searchQuery string
		expected    []TitleCount
	}{
		{
			name:        "Empty search query should return all titles with counts",
			searchQuery: "",
			expected:    []TitleCount{{Title: "Title 1", Count: 2}, {Title: "Title 2", Count: 1}, {Title: "Another", Count: 1}},
		},
		{
			name:        "Search query 'Title' should return titles containing 'Title' with counts",
			searchQuery: "Title",
			expected:    []TitleCount{{Title: "Title 1", Count: 2}, {Title: "Title 2", Count: 1}},
		},
		{
			name:        "Search query 'Another' should return titles containing 'Another' with counts",
			searchQuery: "Another",
			expected:    []TitleCount{{Title: "Another", Count: 1}},
		},
		{
			name:        "Search query 'Unknown' should return empty result",
			searchQuery: "Unknown",
			expected:    []TitleCount{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ListTitleWithPersonCount(titles, credits, tc.searchQuery)
			if !titleCountSliceEqual(result, tc.expected) {
				t.Errorf("Test case '%s' failed: expected %v but got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// Helper function to compare slices of TitleCount
func titleCountSliceEqual(a, b []TitleCount) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Title != b[i].Title || a[i].Count != b[i].Count {
			return false
		}
	}
	return true
}
