package belajargolanggenerics

import (
	"fmt"
	"testing"
)

// Selain generic di function terdapat juga generic di type declaration
type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestString(t *testing.T) {
	names := Bag[string]{"Fizonenda", "Poca", "Sondaya"}
	PrintBag(names)
}

func TestInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(numbers)
}
