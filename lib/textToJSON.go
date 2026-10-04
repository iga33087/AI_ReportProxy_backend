package lib

import (
    "encoding/json"
)

func TextToJSON(text string) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func TextToJSONArray(text string) ([]map[string]any, error) {
	var result []map[string]any
	if err := json.Unmarshal([]byte(text), &result); err != nil {
		return nil, err
	}
	return result, nil
}