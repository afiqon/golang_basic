package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	a = int(b)
	fmt.Println(a, b)

	// var x int
	// x = "5"

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)

	fmt.Printf("Your age is %d\n", 21)
	fmt.Printf("x is %d, y is %f\n", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	//%q for quoted string
	fmt.Printf("This is a %q", figure)

	//%v -> replaced by any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	// %T -> type
	fmt.Printf("figure is of type %T \n", figure)
	fmt.Printf("radius is of type %T \n", radius)
	fmt.Printf("pi is of type %T \n", pi)

	// %t -> bool
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b -> base 2
	fmt.Printf("%b \n", 55)
	fmt.Printf("%08b \n", 55)

	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %.3f\n", x*y)
}
