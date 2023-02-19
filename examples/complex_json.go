package examples

import (
	"encoding/json"
	"log"
	"os"
)

type (
	FullPerson struct {
		Name    string
		Age     int
		Address Address
		Pets    []Pet
	}

	Pet struct {
		Name  string
		Kind  string
		Age   int
		Color string
	}

	Address struct {
		Line1  string
		Line2  string
		Postal string
	}
)

func ComplexJson() {
	b, err := os.ReadFile("assets/complex.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var person FullPerson
	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%v\n", person)
}
