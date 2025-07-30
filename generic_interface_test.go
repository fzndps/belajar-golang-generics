package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Generic juga bisa kita gunakan di Interface secara
// otomatis, semua struct yang ingin mengikuti kontrak interface
// tersebut harus menggunakan generic juga

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

// Implemen struct
type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{Value: "Fizo"}
	fmt.Println(myData)

	result := ChangeValue(&myData, "Dea")
	fmt.Println(result)

	assert.Equal(t, "Dea", result)
	assert.Equal(t, "Dea", myData.Value)

	fmt.Println(myData)

}
