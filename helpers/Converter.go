package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// StringToInt converter
func StringToInt(v string) int {

	r, e := strconv.Atoi(v)
	if e != nil {
		return 0
	}

	return r
}

// MarshalJOSN readable data to bye
func MarshalJOSN(v interface{}) []byte {
	data, _ := json.MarshalIndent(v, "", " ")

	return data
}

// UnmarshalJSON json byte data
func UnmarshalJSON(data []byte, v interface{}) error {

	return json.Unmarshal(data, v)
}

// SliceJoin integer to string
func SliceJoin(v interface{}, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(v), " ", delimiter, -1), "[]")
}
