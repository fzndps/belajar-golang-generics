package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Struct juga mendukung generic, dengan menggunakan generic,
// kita bisa membuat Field dengan tipe data yang sesuai dengan Type Parameter
type Data[T any] struct {
	First  T
	Second T
}

// Selain di function, kita juga bisa tambahkan generic di method
// (function di struct)

// Namun, generic di method merupakan generic yang terdapat di struct nya.

// Kita wajib menyebutkan semua type parameter yang terdapat di struct,
// walaupun tidak kita gunakan misalnya, atau jika tidak ingin kita gunakan,
// kita bisa gunakan _ (garis bawah) sebagai pengganti type parameter nya

// Method tidak bisa memiliki type parameter yang mirip dengan di function

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func (d *Data[T]) ChangeSecond(second T) T {
	d.Second = second
	return d.Second
}

func TestGenericStruct(t *testing.T) {
	myData := Data[string]{
		First:  "Fizonenda",
		Second: "Dea",
	}

	fmt.Println(myData)
}

func SayHelloDawg(name string)  {
	fmt.Println("Sap dawgg "+ name)
}

func TestGenericMethod(t *testing.T) {
	myData := Data[string]{
		First:  "Fizonenda",
		Second: "Dea",
	}

	SayHelloDawg("hallo jawa")

	assert.Equal(t, "Sondaya", myData.ChangeFirst("Sondaya"))
	assert.Equal(t, "Poca", myData.ChangeSecond("Poca"))
	assert.Equal(t, "Hello Deyya", myData.SayHello("Deyya"))
}
