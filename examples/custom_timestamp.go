package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type (
	BirthDate time.Time

	Baby struct {
		Name      string    `json:"name" validate:"required"`
		Gender    string    `json:"gender" validate:"required"`
		BirthDate BirthDate `json:"birth_date" validate:"required"`
	}
)

func (bd BirthDate) String() string {
	return time.Time(bd).Format("02-01-2006")
}

func (bd *BirthDate) UnmarshalJSON(input []byte) error {
	value := strings.Trim(string(input), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("02-01-2006", value)
	if err != nil {
		return err
	}

	*bd = BirthDate(t)
	return nil
}

func (bd *BirthDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(*bd).Format("02-01-2006"))), nil
}

func CustomTimestamp() {
	var baby Baby
	input := []byte(`{"name": "Mary", "gender": "F", "birth_date": "19-02-2023"}`)
	if err := json.Unmarshal(input, &baby); err != nil {
		log.Fatalf("Unable to unmarshal due to %s\n", err)
	}

	log.Printf("Baby %s (%s) was born on %v\n", baby.Name, baby.Gender, baby.BirthDate)

	if err := os.WriteFile("assets/baby.json", input, 0666); err != nil {
		log.Fatalf("Unable to save file due to %s\n", err)
	}
}
