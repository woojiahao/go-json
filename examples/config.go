package examples

import (
	"github.com/knadh/koanf/parsers/json"
	file "github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

var K = koanf.New(".")

func ReadConfig() {
	if err := K.Load(file.Provider("assets/config.json"), json.Parser()); err != nil {
		log.Fatalf("Failed to read configuration due to %s\n", err)
	}

	log.Printf(
		"Title \"%s\" is written by %s and it is described to be\n\t%s\nIt contains %0.f chapters\n",
		K.Get("title"),
		K.Get("author"),
		K.Get("description"),
		K.Get("chapter_count"),
	)
}
