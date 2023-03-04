package examples

import (
	"encoding/json"
	"github.com/sanity-io/litter"
	"log"
	"os"
)

func Decoder() {
	coffeeFile, err := os.Open("assets/coffee.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var coffee UniversalDog
	decoder := json.NewDecoder(coffeeFile)
	err = decoder.Decode(&coffee)
	if err != nil {
		log.Fatalf("Unable to decode due to %s\n", err)
	}

	litter.Dump(coffee)
}

func Encoder() {
	newDog := UniversalDog{
		Breed:         "Poodle",
		Age:           15,
		Name:          "Chloe",
		FavoriteTreat: "Dog Sticks",
	}

	newDogFile, err := os.Create("assets/new_dog.json")
	if err != nil {
		log.Fatalf("Unable to create file due to %s\n", err)
	}

	encoder := json.NewEncoder(newDogFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(newDog)
	if err != nil {
		log.Fatalf("Unable to encode due to %s\n", err)
	}
}
