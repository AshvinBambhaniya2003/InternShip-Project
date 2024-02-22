package models

import (
	"os"
	"reflect"
	"testing"
)

func TestReadCredits(t *testing.T) {
	// Create a temporary CSV file
	file, err := os.CreateTemp("", "test_credits.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write test data to the CSV file
	data := `PersonID,TitleID,Name,Character,Role
7755,tm82169,Bill Baldwin,Fight Announcer,ACTOR
3180,tm82189,Sylvester Stallone,"Robert ""Rocky"" Balboa",ACTOR
7782,tm82115,John G. Avildsen,,DIRECTOR
`
	_, err = file.WriteString(data)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Test reading credits from the created CSV file
	credits, err := ReadCredits(file.Name())
	if err != nil {
		t.Errorf("Error reading credits: %v", err)
	}

	// Expected credits
	expected := []Credit{
		{PersonID: 7755, TitleID: "tm82169", Name: "Bill Baldwin", Character: "Fight Announcer", Role: "ACTOR"},
		{PersonID: 3180, TitleID: "tm82189", Name: "Sylvester Stallone", Character: "Robert Rocky Balboa", Role: "ACTOR"},
		{PersonID: 7782, TitleID: "tm82115", Name: "John G. Avildsen", Character: "", Role: "DIRECTOR"},
	}

	// Compare actual and expected credits
	if !reflect.DeepEqual(credits, expected) {
		t.Errorf("ReadCredits() returned unexpected credits:\nExpected: %v\nActual: %v", expected, credits)
	}
}

func TestGetMostWorkingActor(t *testing.T) {
	// Sample credits
	credits := []Credit{
		{PersonID: 1, TitleID: "tm82169", Name: "Sylvester Stallone", Role: "ACTOR"},
		{PersonID: 2, TitleID: "tm82169", Name: "Sylvester Stallone", Role: "ACTOR"},
		{PersonID: 3, TitleID: "tm82169", Name: "Talia Shire", Role: "DIRECTOR"},
		{PersonID: 4, TitleID: "tm82169", Name: "Talia Shire", Role: "DIRECTOR"},
		{PersonID: 5, TitleID: "tm82169", Name: "Sylvester Stallone", Role: "ACTOR"},
		{PersonID: 6, TitleID: "tm82169", Name: "Talia Shire", Role: "DIRECTOR"},
	}

	result := GetMostWorkingActor(credits)

	// Expected output
	expected := "Sylvester Stallone"

	if result != expected {
		t.Errorf("GetMostWorkingActor() returned unexpected result:\nExpected: %s\nActual: %s", expected, result)
	}
}
