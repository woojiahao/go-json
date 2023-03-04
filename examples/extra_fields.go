package examples

import (
	"encoding/json"
	"github.com/sanity-io/litter"
	"log"
)

func ExtraFields() {
	input := `{
		"Breed": "Golden Retriever", 
		"Age": 8, 
		"Name": "Paws", 
		"FavoriteTreat": "Kibble", 
		"Dislikes": "Cats"
	}`
	var dog Dog
	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	litter.Dump(dog)
}
