package examples

import (
	"encoding/json"
	"log"
)

func SymbolSensitivity() {
	input := `{"Breed": "Golden Retriever", "Age": 8, "Name": "Paws", "favorite_treat": "Kibble"}`
	var dog Dog
	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%s is a %d years old %s who likes %s\n", dog.Name, dog.Age, dog.Breed, dog.FavoriteTreat)
}
