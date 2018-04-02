package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var rectLen, rectWidth float64 = 6, 7

func init() {
	fmt.Printf("Initializing geometry \n")
	if rectLen < 0 {
		log.Fatal("Length cannot be negative")
	}
	if rectWidth < 0 {
		log.Fatal("Width cannot be negative")
	}
}
func main() {

	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
}
