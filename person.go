package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Person struct is contains info about a person
type Person struct {
	FirstName string
	Age       int
}

// ReadPeople converts a JSON file containing people and ages to a human-readable txt file
func ReadPeople() string {

	var people []Person
	var output string

	data, err := ioutil.ReadFile("./testdata/people.json")
	if err != nil {
		fmt.Println("error:", err)
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, p := range people {
		output += fmt.Sprintf("%s is %d years old.\n", p.FirstName, p.Age)
	}
	return output
}
