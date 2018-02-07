package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphaNumClean(t *testing.T) {
	t.Parallel()

	t.Run("AlphaNumClean", func(t *testing.T) {
		x := "1234-aBc*&"

		expected := AlphaNumClean(x)

		actual := "1234aBc"

		assert.Equal(t, expected, actual)
	})
}

func TestParseInteger(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {

		x := "12pspsp*^&^%%$##"

		expected, err := ParseInteger(x)

		actual := int64(12)

		assert.NoError(t, err)

		assert.Equal(t, expected, actual)
	})

}

func TestPathExist(t *testing.T) {
	t.Parallel()

	t.Run("want got true", func(t *testing.T) {

		expected := PathExist(".")

		assert.Equal(t, expected, true)

	})

	t.Run("want got false", func(t *testing.T) {

		expected := PathExist("lhhajja/asasas")

		assert.Equal(t, expected, false)

	})
}

func TestNumericClean(t *testing.T) {
	t.Run("Test numeric clean", func(t *testing.T) {

		r := NumericClean("jag666")

		assert.Equal(t, r, "666")

	})
}
