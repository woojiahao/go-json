package examples

import (
	"encoding/json"
	"github.com/sanity-io/litter"
	"log"
	"os"
)

type BookConfiguration struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Description  string `json:"description"`
	ChapterCount int    `json:"chapter_count"`
}

var Configuration BookConfiguration

func ReadConfig() {
	b, err := os.ReadFile("assets/config.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	err = json.Unmarshal(b, &Configuration)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	litter.Dump(Configuration)
}
