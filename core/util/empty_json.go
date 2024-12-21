package util

import "encoding/json"

func EmptyJsonBracketsBytes() ([]byte, error) {
	return json.Marshal(struct{}{})
}
