package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tipe parameter bisa dibuat untuk parameter, return value, variable
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var resultString string = Length("Fizo")
	assert.Equal(t, "Fizo", resultString)

	var resultInt int = Length(100)
	assert.Equal(t, 100, resultInt)
}
