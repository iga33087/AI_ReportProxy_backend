package lib

import (
    "encoding/json"
)

func TextToJSON[T any](text string) (T, error) {
	var result T
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		return result, err
	}
	return result, nil
}