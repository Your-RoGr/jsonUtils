package jsonUtils

import (
	"encoding/json"
	"errors"
	"regexp"
)

func FindAndParseJSON(text string) (*[]map[string]interface{}, error) {

	jsonRegex := regexp.MustCompile(`\{(?:[^{}]|(?:\{(?:[^{}]|(?:\{[^{}]*\}))*\}))*\}`)

	matches := jsonRegex.FindAllString(text, -1)

	if len(matches) == 0 {
		return nil, errors.New("no JSON found in the text")
	}

	var results []map[string]interface{}

	for _, match := range matches {
		var parsed map[string]interface{}
		err := json.Unmarshal([]byte(match), &parsed)
		if err != nil {
			continue // Skip invalid JSON
		}
		results = append(results, parsed)
	}

	if len(results) == 0 {
		return nil, errors.New("no valid JSON found in the text")
	}

	return &results, nil
}
