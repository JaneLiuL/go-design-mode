package factory

import "testing"

func TestDog(t *testing.T)  {
	animal := NewAnimal("miao")
	result := animal.Animal("qq")
	if result == "miao qq" {
		t.Fatal("new miao wrong")
	}
}
