package helpers

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	t.Parallel()

	t.Run("run for success", func(t *testing.T){
		r := StringToInt("1")
		assert.Equal(t, r, 1)
	})

	t.Run("run for success", func(t *testing.T){
		r := StringToInt("a")
		assert.Equal(t, r, 0)
	})

}

func TestSliceJoin(t *testing.T) {
	t.Parallel()

	id := []int{1, 2, 3}

	r := SliceJoin(id, ",")

	assert.Equal(t, r, "1,2,3")
}

func TestMarshalJOSN(t *testing.T) {
	t.Parallel()

	x := struct {
		A string `json:"a"`
		B string `json:"b"`
	}{
		A: "1",
		B: "2",
	}

	r, err := json.MarshalIndent(x, "", " ")

	assert.NoError(t, err)

	j := MarshalJOSN(x)

	assert.Equal(t, j, r)
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var z, x struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	z.A = "1"
	z.B = "2"

	r := MarshalJOSN(z)

	err := UnmarshalJSON(r, &x)

	assert.NoError(t, err)
}
