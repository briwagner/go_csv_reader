package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type person struct {
	Name  string `json:"name"`
	Dob   string `json:"dob"`
	Color string `json:"color"`
}

// Add a method; using the struct as 'receiver'
func (p *person) sayName() string {
	return p.Name
}

func main() {
	b, err := ioutil.ReadFile("classStats")

	r := csv.NewReader(strings.NewReader(string(b)))
	r.Comma = '|'
	records, err := r.ReadAll()

	totalRecords := len(records)
	fmt.Printf("%v records", totalRecords)

	// How to dynamically set the size here?
	peoples := make([]person, totalRecords)

	if err != nil {
		fmt.Println("error reading file")
	}

	for i, rec := range records {
		if i > totalRecords {
			break
		}
		person := person{rec[0], rec[1], rec[2]}
		peoples[i] = person
	}

	// Specify indent size.
	data, _ := json.MarshalIndent(peoples, "", "    ")
	writeFile("peoples.json", data)

	dataRaw, _ := json.Marshal(peoples)
	writeFile("peoplesRaw.json", dataRaw)
}

func writeFile(filename string, d []byte) {
	err := ioutil.WriteFile(filename, d, 0755)
	if err != nil {
		fmt.Println("Error writing file")
	}
}
