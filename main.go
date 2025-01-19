package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	printff()
}

func DeclaringVariable() {
	var greeting string = "Hello World"
	fmt.Println(greeting)
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
