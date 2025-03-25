package jsonUtils

import (
	"encoding/json"
	"errors"
	"fmt"
)

type numberedJSON struct {
	items      *[]string
	key        *string
	generalKey *string
}

func (nj *numberedJSON) GetString() (string, error) {
	if nj.items == nil {
		return "", errors.New("items cannot be nil")
	}

	result := make(map[string]map[string]string)
	innerMap := make(map[string]string)

	for i, item := range *nj.items {
		innerMap[fmt.Sprintf("%s%d", *nj.key, i+1)] = item
	}

	result[*nj.generalKey] = innerMap

	jsonData, _ := json.Marshal(result)

	return string(jsonData), nil
}

func (nj *numberedJSON) GetMap() *map[string]map[string]string {
	if nj.items == nil {
		return &map[string]map[string]string{
			"General": {},
		}
	}

	result := make(map[string]map[string]string)
	innerMap := make(map[string]string)

	for i, item := range *nj.items {
		innerMap[fmt.Sprintf("%s%d", *nj.key, i+1)] = item
	}

	result[*nj.generalKey] = innerMap

	return &result
}

func CreateNumberedJSON(items *[]string, key string, generalKey string) *numberedJSON {
	return &numberedJSON{
		items:      items,
		key:        &key,
		generalKey: &generalKey,
	}
}
