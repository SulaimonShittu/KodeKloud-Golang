package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var course string = "Computer Science"

func main() {
	convtype()
}

func convtype() {
	//Converting from float to int & vice versa
	var cgpa float64 = 3.54
	var fgpa = int(cgpa)
	fmt.Printf("Your GP is : %.4f, rounded to %d\n", cgpa, fgpa)

	//Converting from strings to integers - using the strconv package
	var age10 string = "20"
	age11, err := strconv.Atoi(age10)
	if err == nil {
		fmt.Printf("The variable age11 is of Type : %T, and Value : %v\n", age11, age11)
	}

	//Converting from integers to strings - using the strconv package
	var age00 int = 21
	age01 := strconv.Itoa(age00)
	fmt.Printf("The variable age01 is of Type : %T, and Value : %v\n", age01, age01)
}

//Knowing Types of Variables & Values

func gettype() {
	var grades int = 80
	var message string = "Excellent"

	// using the %T - format specifier
	fmt.Printf("The variable grades is of Type : %T\n", grades)
	fmt.Printf("The variable message is of Type : %T\n", message)
	fmt.Printf("The value 68.44 is of Type : %T\n", 68.44)

	//using the reflect.TypeOf Method
	fmt.Printf("The variable grades is of Type : %v\n", reflect.TypeOf(grades))
	fmt.Printf("The variable message is of Type : %v\n", reflect.TypeOf(message))
	fmt.Printf("The value 68.44 is of Type : %v\n", reflect.TypeOf(68.44))
}

// User Input
func userinput() {
	var name string
	var isstudent bool
	count, err := fmt.Scanf("%s %t", &name, &isstudent)
	if err == nil && count == 2 && isstudent == true {
		fmt.Printf("Welcome, student : %s", name)
	}
}

// Zero Values
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

// Global Variables
func glovar() {
	fmt.Println(course)
	course = "CSC"
	fmt.Println(course)
}

// Declaring Variables
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

// PrintF function & format specifiers
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
