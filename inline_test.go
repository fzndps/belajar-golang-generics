package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sebelum-sebelumnya, kita selalu menggunakan type declaration
// atau type set ketika membuat type constraint di type parameter

// Sebenarnya tidak ada kewajiban kita harus membuat type
// declaration atau type set jika kita ingin membuat type
// parameter, kita bisa gunakan secara langsung (in line)
// pada type constraint, misalnya di awal kita sudah bahas tentang
// interface {} (kosong), tapi kita selalu gunakan type
// declaration any

// Jika kita mau, kita juga bisa langsung gunakan
// interface { int | float32 | float64} dibanding membuat
// type set Number misalnya

func FindMax[T interface{ int | int64 | float64 }](first, second T) T {
	if first > second {
		return first
	} else {
		return second
	}
}

func TestInline(t *testing.T) {
	assert.Equal(t, 100, FindMax(50, 100))
	assert.Equal(t, int64(100), FindMax(int64(100), int64(50)))
	assert.Equal(t, float64(100.0), FindMax(100.0, 50.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	Numbers := []int{
		1, 2, 3, 4, 5, 6,
	}

	first := GetFirst[[]int, int](Numbers)
	assert.Equal(t, 1, first)
}
