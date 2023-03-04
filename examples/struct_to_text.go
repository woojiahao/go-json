package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

func StructToText() {
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
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b))
}
