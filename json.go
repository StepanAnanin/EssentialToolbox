package essential

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

// Decode given json.
// Returns decoded json and true if there are no errors, false otherwise.
func DecodeJSON[T any](input io.Reader) (T, bool) {
	var result T

	if err := json.NewDecoder(input).Decode(&result); err != nil {
		log.Printf("\n[ ERROR ] Failed to decode JSON\n%s", err)

		return result, false
	}

	return result, true
}

func DecodeStringJSON[T any](input string) (T, bool) {
	return DecodeJSON[T](strings.NewReader(input))
}

// Encode passed `target` argument into json.
// Returns `target` argument in json format ([]byte), and true if no errors occured, false otherwise.
func EncodeJSON(target interface{}) ([]byte, bool) {
	result, err := json.Marshal(target)

	if err != nil {
		log.Printf("\n[ ERROR ] Failed to marshal json\n%s", err)

		return nil, false
	}

	return result, true
}
