package examples

import (
	"encoding/json"
	"log"
)

type (
	TypeAliasExample string

	TypeAliasStruct struct {
		Example TypeAliasExample
	}
)

func TypeAlias() {
	input := `{"Example": "Hello world"}`
	var tas TypeAliasStruct
	err := json.Unmarshal([]byte(input), &tas)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	log.Printf("%v\n", tas)
}
