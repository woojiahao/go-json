package examples

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Email   string
	Phone   string
	Hobbies []string
}

func FileToStruct() {
	b, err := os.ReadFile("assets/sample.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var person Person
	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%v\n", person)
}
