package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

// Kita bisa memasukkan sebuah constraint Employee atau sebuah interface ke dalam function generics
// Artinya kita bisa memasukkan apapun itu ke parameter selama kompatibel dengan Employee
func GetName[T Employee](parameter T) string {
	// Kalau type parameter generic ditentukan sama dengan interface maka
	// dapat mengakses semua method yang ada di interface yang digunakan
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresident() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresident() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Fizo", GetName[Manager](&MyManager{Name: "Fizo"}))
	assert.Equal(t, "Fizo", GetName[VicePresident](&MyVicePresident{Name: "Fizo"}))
}
