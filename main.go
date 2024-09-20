package faker

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed data/first_names.csv
var firstNamesData string

//go:embed data/last_names.csv
var lastNamesData string

var firstNames []string
var lastNames []string

func init() {
	firstNames = parseCSV(firstNamesData)
	lastNames = parseCSV(lastNamesData)
}

func parseCSV(data string) []string {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("failed to parse CSV data: %v", err))
	}

	var names []string
	for _, record := range records {
		names = append(names, record[0])
	}
	return names
}

// GenerateFirstName returns a random first name.
func GenerateFirstName() string {
	return getRandomName(firstNames)
}

// GenerateLastName returns a random last name.
func GenerateLastName() string {
	return getRandomName(lastNames)
}

// GenerateFullName returns a random full name.
func GenerateFullName() string {
	return fmt.Sprintf("%s %s", GenerateFirstName(), GenerateLastName())
}

func getRandomName(names []string) string {
	return names[rand.Intn(len(names))]
}
