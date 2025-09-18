package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// Initialization Section
////////////////////////////////////////////////////////////////////////////////

// init function is a special function in Go.
// It automatically executes **before the main() function**.
// 
// Common use cases:
// - Initialize global variables
// - Setup configurations
// - Preload resources
//
// NOTE:
// - You can have multiple `init()` functions within the same package.
// - They run in the order they are defined.
func init() {
	fmt.Println("Init: This should be the first to execute")
}

////////////////////////////////////////////////////////////////////////////////
// Condition Function
////////////////////////////////////////////////////////////////////////////////

// Condition demonstrates basic user input handling and conditional statements.
// It asks for the user's age and checks if they are allowed to perform an action.
func Condition() {
	var age int
	fmt.Println("Please Input Your Age:")
	fmt.Scanln(&age) // Reads user input from the console

	// Simple if-else conditional check
	if age < 18 {
		fmt.Println("Oops, you cannot do this")
	} else {
		fmt.Println("Yay, you can do this!")
	}
}

////////////////////////////////////////////////////////////////////////////////
// Structs and Methods
////////////////////////////////////////////////////////////////////////////////

// User represents a basic user with a name and age.
type User struct {
	name string
	age  int
}

// Global variable to demonstrate how to initialize a struct at the package level.
var user1 = User{name: "Sojib", age: 24}

// getName returns the user's name.
// It is a value receiver method, meaning it works on a copy of the struct.
func (user User) getName() string {
	return user.name
}

// getAge returns the user's age.
// It is also a value receiver method.
func (user User) getAge() int {
	return user.age
}

////////////////////////////////////////////////////////////////////////////////
// Arrays and Pointers
////////////////////////////////////////////////////////////////////////////////

// arr is a fixed-size array of 5 integers.
// NOTE: Arrays in Go are value types and have a fixed length.
var arr = [5]int{1, 2, 3, 4, 5}

// getUserPointer demonstrates how to work with pointers in Go.
// It takes a pointer to an array and returns the same pointer.
//
// Example usage:
// arrPtr := getUserPointer(&arr)
func getUserPointer(usr *[5]int) *[5]int {
	return usr
}

////////////////////////////////////////////////////////////////////////////////
// Higher-Order Functions
////////////////////////////////////////////////////////////////////////////////

// higherOrderFunc is a higher-order function that accepts another function
// as an argument and executes it with two integer values.
//
// Parameters:
// - HOF: a function that takes two integers and returns nothing.
//
// Example usage:
// higherOrderFunc(sumFunc)
func higherOrderFunc(HOF func(int, int)) {
	HOF(10, 20)
}

// sumFunc is a simple function that adds two integers and prints the result.
func sumFunc(num1 int, num2 int) {
	fmt.Println("The Sum Is:", num1+num2)
}

////////////////////////////////////////////////////////////////////////////////
// Utility Functions
////////////////////////////////////////////////////////////////////////////////

// getFullDetails prints the complete details of a User.
// This demonstrates passing structs by value to a function.
func getFullDetails(usr User) {
	fmt.Println("The Name is from full details:", usr.name)
	fmt.Println("The Age is from full details:", usr.age)
}

////////////////////////////////////////////////////////////////////////////////
// Main Function
////////////////////////////////////////////////////////////////////////////////

// main is the entry point of the Go program.
func main() {
	fmt.Println("Welcome To Golang")

	// Example of using a higher-order function
	higherOrderFunc(sumFunc)

	// Anonymous function example
	func(name string) {
		fmt.Println("Hello", name)
	}("Sojib")

	// Printing user details using a helper function
	getFullDetails(user1)

	// Demonstrating working with pointers
	usrDetails := getUserPointer(&arr)
	fmt.Println("The User Details are:", usrDetails)

	// Accessing struct methods
	fmt.Println("The name is:", user1.getName())

	// Printing a fixed-size array
	fmt.Println(arr)

	// Demonstrating a slice, which is more flexible than arrays
	arr2 := []int{6, 7, 8, 9, 10}
	fmt.Println(arr2)

	// Run the condition checker for user input
	Condition()
}
