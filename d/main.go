package main

import (
	"fmt"
	"log"
)

// greet interface
type greet interface {
	Greet() string
}

// impl Greet in vietnamese
type vietnamese struct{}

func (v vietnamese) Greet() string {
	return fmt.Sprintf("%T: Xin chào!", v)
}

// impl Greet in english
type english struct{}

func (e english) Greet() string {
	return fmt.Sprintf("%T: Hello!", e)
}

func greeter (g greet) {
	log.Println(g.Greet())
}


func main() {
	x := []greet{
		english{},
		vietnamese{},
	}

	for _, e := range x {
		greeter(e)
	}

}
