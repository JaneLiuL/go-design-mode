package factory

// This is abstract factory mode example
// 抽象工厂一般是用于生产产品的时候，被生成的产品是有关系的

import "fmt"

type AnimalFactory interface {
	CreateAnimal() Animal
}
type Animal interface {
	Eat()
}

type DogFactory struct {}
type CatFactory struct {}
type Dog struct {}
type Cat struct {}

// 狗工厂，生产狗
func (dog DogFactory) CreateAnimal() Animal {
	return &Dog{}
}

// 猫工厂，生产猫
func (cat CatFactory) CreateAnimal() Animal {
	return &Cat{}
}

func (dog Dog) Eat()  {
	fmt.Println("dog eat meat")
}

func (cat Cat) Eat()  {
	fmt.Println("cat eat milk")
}

func ProductAnimalAndEat()  {
	var factory AnimalFactory
	factory = &DogFactory{}
	dog := factory.CreateAnimal()
	dog.Eat()

	factory = &CatFactory{}
	cat := factory.CreateAnimal()
	cat.Eat()
}