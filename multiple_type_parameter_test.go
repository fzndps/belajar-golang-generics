package belajargolanggenerics

import (
	"fmt"
	"testing"
)

// Type parameter bisa 2 atau lebih cukup pisahkan dengan koma dan bedakan namanya
func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter("test", 2)
	MultipleParameter(100, "test2")
}
