package examples

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
)

type ValidatedUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,min=18,max=99"`
}

func ValidateBadUser() {
	b, err := os.ReadFile("assets/validate_user_bad.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var user ValidatedUser
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	log.Printf("User before validation: %v\n", user)

	err = validator.New().Struct(user)
	if err != nil {
		log.Fatalf("Validation failed due to %v\n", err)
	}
}

func ValidateGoodUser() {
	b, err := os.ReadFile("assets/validate_user_good.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	var user ValidatedUser
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	log.Printf("User before validation: %v\n", user)

	err = validator.New().Struct(user)
	if err != nil {
		log.Fatalf("Validation failed due to %v\n", err)
	}

	log.Printf("User after validation: %v\n", user)
}
