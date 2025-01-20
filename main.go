package main

import "fmt"

var course string = "Computer Science"

func main() {
	userinput()
}

func userinput() {
	var name string
	var isstudent bool
	count, err := fmt.Scanf("%s %t", &name, &isstudent)
	if err == nil && count == 2 && isstudent == true {
		fmt.Printf("Welcome, student : %s", name)
	}
}

func ZeroValues() {
	var (
		a int
		b string
		c rune
		d float64
		e bool
	)
	fmt.Printf("The Zero Values of types,\nint(%T) : %d,\nstring(%T) : %s,\nrune(%T) : %s,\nfloat64(%T) : %f,\nbool(%T) : %t", a, a, b, b, c, string(c), d, d, e, e)
}

func glovar() {
	fmt.Println(course)
	course = "CSC"
	fmt.Println(course)
}

func DeclaringVariable() {
	var greeting string = "Hello World"
	fmt.Println(greeting)

	// Variables of same type dclaration
	var a, b int = 20, 45
	// Varaibles of different types declaration
	fmt.Println("a : ", a, " ", "b ; ", b)
	var (
		name string = "Sulaimon"
		age  int    = 21
	)
	fmt.Printf("Your name is : %s, and age : %d", name, age)

}

func printff() {
	// format specifiers
	// %v - formats value in a default format
	// %T - type of the value
	// %c -character
	// %q - quoted character or string
	// %d - used with decimals
	// %s - plain string
	// %f - floating points
	// %t - boolean
	var name string = "Sula"
	var score int = 87
	fmt.Printf("Hey %v, %q you have scored %d/100 in your exam", name, "HeadOfTheTable", score)
}
