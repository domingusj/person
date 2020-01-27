package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Person struct is contains info about a person
type Person struct {
	FirstName string
	Age       int
}

// ReadPeople converts a JSON file containing people and ages to a human-readable txt file
func ReadPeople() {

	var people []Person

	data, err := ioutil.ReadFile("./testdata/people.json")
	if err != nil {
		fmt.Println("error:", err)
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i := range people {
		fmt.Println(people[i].FirstName + " is " + strconv.Itoa(people[i].Age) + " years old.")
	}

}
