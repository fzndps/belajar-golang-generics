package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AgeInference int

type NumberInference interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func MinInference[T NumberInference](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

// Type Inference merupakan fitur dimana kita tidak perlu
// menyebutkan Type Parameter ketika memanggil kode generic
// (kurung kotak [])

// Tipe data Type Parameter bisa dibaca secara otomatis misal
// dari parameter yang kita kirim

// Namun perlu diingat, pada beberapa kasus, jika terjadi error
// karena Type Inference, kita bisa dengan mudah memperbaikinya
// dengan cara menyebutkan Type Parameter nya saja

func TestInference(t *testing.T) {
	assert.Equal(t, 50, MinInference(100, 50))
	assert.Equal(t, int64(100), MinInference(int64(100), int64(200)))
	assert.Equal(t, float64(100), MinInference(float64(100), float64(200)))
	assert.Equal(t, Age(100), MinInference(Age(100), Age(200)))
}
