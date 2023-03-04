package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Breed         string
	Age           int
	Name          string
	FavoriteTreat string
}

func TextToStruct() {
	input := `{
		"Breed": "Golden Retriever", 
		"Age": 8, 
		"Name": "Paws", 
		"FavoriteTreat": "Kibble"
	}`
	var dog Dog
	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	fmt.Printf("%s is a %d years old %s who likes %s\n", dog.Name, dog.Age, dog.Breed, dog.FavoriteTreat)
}
