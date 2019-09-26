package mapper

import (
	"encoding/json"
	"fmt"
)

func ToJson(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ToObject(jsonStr string, v interface{}) error {
	bytes := []byte(jsonStr)
	err := json.Unmarshal(bytes, v)
	return err
}

func MapTo(origin, target interface{}) error {
	bytes, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, target)
	if err != nil {
		return err
	}

	return nil
}
