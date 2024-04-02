package main

import "log"

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

// square only need to impl shape.area()
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

// cube should impl both
type cube struct {
	length float64
}

func (c cube) area() float64 {
	return 6 * c.length * c.length
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

// sum area
func areaSum(shapes ...shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// sum of area and volume
func areaVolumeSum(shapes ...object) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s := square{5}
	c1 := cube{10}
	c2 := cube{20}

	log.Printf("areaSum %f", areaSum(s, c1, c2))
	log.Printf("areaVolumeSum %f", areaVolumeSum(c1, c2))
}
