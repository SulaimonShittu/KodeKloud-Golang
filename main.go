package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var course string = "Computer Science"

func main() {
	maps()
}

// functions

//variadicfunctions

// named returns
func operation(a int, b int) (sum int, dif int) {
	sum, dif = a+b, a-b
	return
}

// function basics
func greetings() {
	fmt.Println("Hey Joe !")
}

// returning a value
func addNumbers(a int, b int) (sum int) {
	sum = a + b
	return
}

func subNumbers(a int, b int) int {
	dif := a - b
	return dif
}

func maps() {
	codes := map[string]string{"En": "English", "Hi": "Hindi", "Fr": "French"}
	fmt.Println(len(codes))

	// accessing a key value
	fmt.Println(codes["Hi"])
	fmt.Println(codes["Fr"])

	// key validity
	english, found := codes["En"]
	if found {
		fmt.Printf("%s is %t\n", english, found)
	}
	french, found := codes["Fr"]
	if found {
		fmt.Printf("%s is %t\n", french, found)
	}

	// updating a value
	codes["Sl"] = "Slovenia"
	codes["Sl"] = "Slovakia"

	// deleting a value
	delete(codes, "Hi")
	fmt.Println(codes)

	// truncating a map
	codes = make(map[string]string)
}

func slices() {
	// declaring & initializing a slice
	pets := []string{"Cats", "Dogs", "Mice"}
	fmt.Println(pets)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)
	slice_2 := arr[0:3]
	fmt.Println(slice_2)

	slice_3 := make([]int, 8, 12)
	for k, _ := range slice_3 {
		slice_3[k] = k * k
	}

	// slice refrencing arrays
	ages := [5]int{19, 20, 21, 22, 23}
	slice_age := ages[1:3]
	fmt.Println(ages)
	fmt.Println(slice_age)
	fmt.Println("After modification")
	slice_age[0] = 42
	fmt.Println(ages)
	fmt.Println(slice_age)
	fmt.Println(len(slice_age))
	fmt.Println(cap(slice_age))

	//appending
	slice_age = append(slice_age, 24, 31, 27)
	fmt.Println(len(slice_age))
	fmt.Println(cap(slice_age))

	//appending two slices
	slice_5 := append(slice_2, slice_3...)
	fmt.Println(slice_5)

	//copying slcies
	slice_6 := []int{5, 9, 2, 7, 4}
	slice_7 := make([]int, 5, 10)
	copy(slice_7, slice_6)
}

func arrays() {
	// array initialization
	var fruits [2]string = [2]string{"Apples", "Oranges"}
	fmt.Println(fruits)

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	// len method
	fmt.Println(len(fruits))
	fmt.Println(len(marks))
	fmt.Println(len(names))

	//array indexing
	var songs [3]string = [3]string{"Genie", "Tell me", "Shotta Soul"}
	songs[0] = "Never Lie"

	for i := 0; i < len(songs); i++ {
		fmt.Println(songs[i])
	}

	for k, v := range songs {
		fmt.Printf("%d => %s\n", k, v)
	}

	//Multi-dimensional arrays
	var scores [3][2]int = [3][2]int{
		{99, 87},
		{89, 97},
		{79, 78},
	}
	fmt.Println(scores[1][0])
}

func forloops() {
	//traditional for loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}

	//while loop styled loops
	a := 1
	for a <= 5 {
		fmt.Println(a * a)
		a++
	}
	//Range-styled for loops
	for i := range 10 {
		fmt.Printf("%d x %d = %d\n", i, i, i*i)
	}

}

func switchoperator() {
	// Switch on variables
	var age int = 19
	switch age {
	case 18:
		fmt.Println("Welcome to the legal age.")
	case 19:
		fmt.Println("You are now a big boy.")
		fallthrough
	case 20, 21, 22:
		fmt.Println("The famous 20's")
		fallthrough
	default:
		fmt.Println("You sha have an age.")
	}

	//switch on conditionals
	_, err := fmt.Scanf("%d", &age)
	if err == nil {
		switch {
		case age < 18:
			fmt.Println("You aren't up to legal age yet.")
		case age >= 18:
			fmt.Println("You can now vote & drink alcohol.")
		}
	}
}

func bitwiseoperators() {
	//bitwise and
	var a, b int = 16, 25
	z := a & b
	fmt.Println(z)

	//bitwise or
	z = a | b
	fmt.Println(z)

	//bitwise xor
	z = a ^ b
	fmt.Println(z)

	//bitwise left-shift
	z = 566 << 6
	fmt.Println(z)

	//bitwise right-shift
	z = 456 >> 4
	fmt.Println(z)

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
	fmt.Println("a : ", a, " ", "b : ", b)
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
	// %d - used with numbers/decimals
	// %s - plain string
	// %f - floating points
	// %t - boolean
	var name string = "Sula"
	var score int = 87
	fmt.Printf("Hey %v, %q you have scored %d/100 in your exam", name, "HeadOfTheTable", score)
}
