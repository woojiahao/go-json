package examples

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func BindBody(body io.ReadCloser, target any) error {
	contents, err := io.ReadAll(body)
	if err != nil {
		log.Fatalf("Failed to read request body")
	}

	err = json.Unmarshal(contents, target)
	if err != nil {
		log.Fatalf("Failed to unmarshal request body")
	}

	return validator.New().Struct(target)
}

type Response map[string]interface{}

func JSON(response Response) []byte {
	j, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal contents")
	}

	return j
}

func StartServer() {
	http.HandleFunc("/request", func(writer http.ResponseWriter, request *http.Request) {
		var login Login
		if err := BindBody(request.Body, &login); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write(JSON(Response{
				"message": "Invalid login request",
			}))
			return
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(JSON(Response{
			"message": fmt.Sprintf("Welcome back %s", login.Username),
		}))
	})
	_ = http.ListenAndServe(":8080", nil)
}
