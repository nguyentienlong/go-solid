package main

import "log"

// liskov substitution principle (LSP
// the objects of super class should be replacable with objects of
// a subclass without affecting the correctness of the program.

// Flying interface
type Flying interface {
	Fly()
}

// Swimming interface
type Swimming interface {
	Swim()
}

// Perfomable interface
type Performable interface {
	Perform()
}

// Sparrow struct
type Sparrow struct {}

func (s Sparrow) Fly() {
	log.Println("sparrow can fly")
}

func (s Sparrow) Perform() {
	s.Fly()
}

// Penguin struct
type Penguin struct {}

func (p Penguin) Swim () {
	log.Println("Penguin can swim")
}

func (p Penguin) Perform() {
	p.Swim()
}

// Animal Show
func AnimalShow(p ...Performable) {
	for _, a := range p {
		a.Perform()
	}
}

func main () {
	s := Sparrow{}
	p := Penguin{}

	AnimalShow(s, p)
}
