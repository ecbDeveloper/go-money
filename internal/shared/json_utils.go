package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJson[T any](w http.ResponseWriter, statusCode int, data T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func DecodeValidJson[T Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, map[string]string{"error": "bad request"}, fmt.Errorf("failed to decode json %w", err)
	}

	problems := data.Valid(r.Context())
	if len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("failed to decode json %w", err)
	}

	return data, nil
}
