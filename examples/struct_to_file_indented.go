package examples

import (
	"encoding/json"
	"log"
	"os"
)

func StructToFileIndented() {
	p := Person{
		Name:  "John Jones",
		Age:   26,
		Email: "johnjones@email.com",
		Phone: "89910119",
		Hobbies: []string{
			"Swimming",
			"Badminton",
		},
	}
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	err = os.WriteFile("assets/person_indented.json", b, 0666)
	if err != nil {
		log.Fatalf("Unable to save to file due to %s\n", err)
	}
}
