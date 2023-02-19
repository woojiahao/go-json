package examples

import (
	"encoding/json"
	"log"
)

func ExtraFields() {
	input := `{"Breed": "Golden Retriever", "Age": 8, "Name": "Paws", "FavoriteTreat": "Kibble", "Dislikes": "Cats"}`
	var dog Dog
	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%v\n", dog)
}
