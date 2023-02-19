package examples

import (
	"encoding/json"
	"log"
)

func MissingFields() {
	input := `{"Breed": "Golden Retriever", "Age": 8, "Name": "Paws"}`
	var dog Dog
	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%s likes %s\n", dog.Name, dog.FavoriteTreat)
}
