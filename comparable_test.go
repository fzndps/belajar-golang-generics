package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1 T, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.False(t, IsSame("fizo", "dea"))
	assert.True(t, IsSame("fizo", "fizo"))
	assert.False(t, IsSame(100, 99))
	assert.True(t, IsSame(100, 100))
}
