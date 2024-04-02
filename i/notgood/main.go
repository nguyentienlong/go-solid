package main

import "log"

// only 1 interface
type shape interface {
	area() float64
	volume() float64
}

// square
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) volume() float64 {
	return 0
}

// cube
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
func areaSum(shapes ...shape) float64{
	sum := 0.0
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// sum of area and volume
func areaVolumeSum(shapes ...shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s := square{5}
	c := cube{10}

	log.Printf("areaSum %f", areaSum(s, c))
	log.Printf("areaVolumeSum %f", areaVolumeSum(s, c))
}
