package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateRandChars, unit testing generate random chars //
func TestCreateRandChars(t *testing.T) {
	var (
		strLen, resultExpected int = 10, 10
	)

	actualResult := CreateRandChars(strLen)
	assert.EqualValues(t, resultExpected, len(actualResult), "Expected result length should be ", resultExpected, " but actual length ", len(actualResult))
}
