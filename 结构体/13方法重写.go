package main 

import "fmt"

type Person struct {
}

func (this *Person) Eat() {
	fmt.Println("Person Eat")
}

func (this *Person) Run() {
	fmt.Println("Person Run")
}

func (this *Person) Sleep() {
	fmt.Println("Person Sleep")
}

type Man struct {
	Person
}

func (this *Man) Eat() {
	// 类似于Java的 super.Eat()
	this.Person.Eat()
	fmt.Println("Man Eat")
}

func (this *Man) Run() {
	fmt.Println("Man Run")
}


func main() {
	m := &Man{}
	m.Eat()
	m.Run()
	m.Sleep()
}

