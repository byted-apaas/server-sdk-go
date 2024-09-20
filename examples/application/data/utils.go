package main

import (
	"encoding/json"
)

func deepCopy(src map[string]interface{}) (map[string]interface{}, error) {
	var dst map[string]interface{}

	bytes, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
