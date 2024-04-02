package main

import "log"

// liskov substitution principle (LSP
// the objects of super class should be replacable with objects of
// a subclass without affecting the correctness of the program.

type Animal interface {
	Fly()
	Swim()
}

type Sparrow struct {}

func (s Sparrow) Fly() {
	log.Println("sparrow can fly")
}

func (s Sparrow) Swim() {
	panic("sparrow can't swim")
}

type Penguin struct {}

func (p Penguin) Fly () {
	panic("Penguin can fly")
}

func (p Penguin) Swim () {
	log.Println("Penguin can swim")
}

func AnimalAction(a Animal) {
	a.Fly()
	a.Swim()
}


func main () {
	s := Sparrow{}
	p := Penguin{}

	AnimalAction(s)
	AnimalAction(p)
}
