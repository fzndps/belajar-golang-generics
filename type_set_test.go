package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Dengan fitur type set ini, kita bisa menentukan lebih dari satu tipe constraint
// yang diperbolehkan pada type parameter
type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Max[T Number](first, second T) T {
	if first > second {
		return first
	} else {
		return second
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, 100, Max(100, 50))
	assert.Equal(t, int64(200), Max(int64(100), int64(200)))
	assert.Equal(t, float64(200), Max(float64(100), float64(200)))
}
