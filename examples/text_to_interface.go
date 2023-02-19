package examples

import (
	"encoding/json"
	"log"
)

func TextToInterface() {
	input := `{"num": 5.1, "arr": ["a", 1, "c"]}`
	var target map[string]interface{}
	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	for k, v := range target {
		log.Printf("k: %s, v: %v", k, v)
	}
}
