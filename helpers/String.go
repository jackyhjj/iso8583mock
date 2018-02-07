// Package helpers implements commonly used functions (string manipulation) //
package helpers

import (
	"math/rand"
	"strings"
	"time"
)

// CreateRandChars -  generate random characters with defined length
// CreateRandChars, generate random characters with defined length //
// input strlen int
// output string
func CreateRandChars(strlen int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var (
		r      *rand.Rand
		result string
	)

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < strlen; i++ {
		index := r.Intn(len(chars))
		result += chars[index: index+1]
	}

	return strings.ToUpper(result) // return upper case result //
}

//RemoveLeftPad - remove string left Padding
func RemoveLeftPad(strlen string, padding string) string{
	isRemove := true
	charString := []rune(strlen)
	result := ""
	for _, y := range charString {
		if string(y) == padding {
			if !isRemove{
				result += string(y)
			}
			continue
		}else{
			isRemove = false
			result += string(y)
		}
	}

	return result
}
