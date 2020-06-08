package main

import "fmt"


type AbstractDog struct {
	Sleep func()
}

func (this *AbstractDog) Eat() {
	fmt.Println("AbstractDog Eat")
	this.Sleep()
}

func (this *AbstractDog) Run() {
	fmt.Println("AbstractDog Run")
}

// 秋田犬
type Akita struct {
	AbstractDog
}

func NewAkita() *Akita {
	ptr := &Akita{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}

func (this *Akita) Sleep() {
	fmt.Println("Akita Sleep")
}

// 拉布拉多犬
type Labrador struct {
	AbstractDog
}

func NewLabrador() *Labrador {
	ptr := &Labrador{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}

func (this *Labrador) Sleep() {
	fmt.Println("Labrador Sleep")
}

func main() {
	akita := NewAkita()
	akita.Eat()

	labrador := NewLabrador()
	labrador.Eat()
}
