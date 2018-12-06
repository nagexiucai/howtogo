package main

import "fmt"

type Animal interface {
	Show()
	Eat()
	Walk()
}
type AnimalLooks struct {
	Eye int
	Ear int
	Nose int
	Mouth int
	Hand int
	Leg int
}
func (x AnimalLooks) Show() {
	fmt.Println("Eye:",x.Eye,"Ear:",x.Ear,"Nose:",x.Nose,"Mouth:",x.Mouth,"Hand:",x.Mouth,"Leg:",x.Leg)
}
func (x AnimalLooks) Eat() {
	fmt.Println("AO~~~")
}
func (x AnimalLooks) Walk() {
	fmt.Println("121")
}

type Dog struct {
	AnimalLooks
	_1212 string
}
func (d Dog) Show() {
	d.AnimalLooks.Show()
	fmt.Println(d._1212)
}
func (d Dog) Cry() {
	fmt.Println("Bark")
}

type Cat struct {
	AnimalLooks
	_1111 string
}
func (c Cat) Show() {
	c.AnimalLooks.Show()
	fmt.Println(c._1111)
}
func (c Cat) Cry() {
	fmt.Println("Miaow")
}

func main() {
	var dog = new(Dog)
	dog.Show()
	dog.Eat()
	dog.Walk()
	dog.Cry()
	var cat = new(Cat)
	cat.Show()
	cat.Eat()
	cat.Walk()
	cat.Cry()
}
