package factory

import "fmt"
// 这个是简单工厂模式， 一般来说Go是使用了Newxx来初始化的，所以Newxx返回接口的时候就是简单工厂模式

type CreateAnimal interface {
	Animal(description string) string
}

type dog struct {}
type cat struct {}

func NewAnimal(description string) CreateAnimal {
	if description == "miao" {
		return &dog{}
	} else {
		return &cat{}
	}

}
func (dog dog) Animal(name string) string {
	return fmt.Sprintf("miao %s" + name)
}
func (cat cat) Animal(name string) string {
	return fmt.Sprintf("wang %s", cat)
}