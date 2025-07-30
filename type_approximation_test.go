package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Kadang, kita sering membuat Type Declaration di Golang untuk
// tipe data lain, misal kita membuat tipe data Age untuk tipe data int

// Secara default, jika kita gunakan Age sebagai type declaration
// untuk int, lalu kita membuat Type Set yang berisi constraint int,
// maka tipe data Age dianggap tidak compatible dengan Type Set yang kita buat

type Age int

// Untungnya, Go-Lang memiliki feature bernama Type Approximation,
// dimana kita bisa menyebutkan bahwa semua constraint dengan tipe tersebut
// dan juga yang memiliki tipe dasarnya adalah tipe tersebut, maka bisa digunakan

// Untuk menggunakan Type Approximation, kita bisa gunakan tanda ~ (tilde)

type NumberApprox interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T NumberApprox](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 50, Min(100, 50))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, float64(100), Min(float64(100), float64(200)))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
