package examples

import (
	"encoding/json"
	"github.com/sanity-io/litter"
	"log"
	"os"
)

type UniversalDog struct {
	Breed         string `json:"breed"`
	Age           int    `json:"age"`
	Name          string `json:"name"`
	FavoriteTreat string `json:"favorite_treat"`
}

func StructTagsUnmarshal() {
	b, err := os.ReadFile("assets/coffee.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var coffee UniversalDog
	err = json.Unmarshal(b, &coffee)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(coffee)
}

func StructTagsMarshal() {
	dog := UniversalDog{
		Name:          "Millie",
		Age:           4,
		Breed:         "Golden Retriever",
		FavoriteTreat: "Ground Beef",
	}
	b, err := json.MarshalIndent(dog, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	err = os.WriteFile("assets/millie.json", b, 0666)
	if err != nil {
		log.Fatalf("Unable to save to file due to %s\n", err)
	}
}

type User struct {
	Username string   `json:"username"`
	Password string   `json:"-"`
	Email    string   `json:"email"`
	Hobbies  []string `json:"hobbies,omitempty"`
}

func StructTagsOthers() {
	user := User{
		Username: "admin",
		Password: "root",
		Email:    "admin@email.com",
	}

	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	err = os.WriteFile("assets/user.json", b, 0666)
	if err != nil {
		log.Fatalf("Unable to save to file due to %s\n", err)
	}
}
