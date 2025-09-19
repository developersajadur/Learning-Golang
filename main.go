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
func (user User) getName() string {
	return user.name
}

// getAge returns the user's age.
func (user User) getAge() int {
	return user.age
}

////////////////////////////////////////////////////////////////////////////////
// Arrays and Pointers
////////////////////////////////////////////////////////////////////////////////

// arr is a fixed-size array of 5 integers.
var arr = [5]int{1, 2, 3, 4, 5}

// getUserPointer demonstrates how to work with pointers in Go.
func getUserPointer(usr *[5]int) *[5]int {
	return usr
}

////////////////////////////////////////////////////////////////////////////////
// Higher-Order Functions
////////////////////////////////////////////////////////////////////////////////

// higherOrderFunc is a higher-order function that accepts another function
// as an argument and executes it with two integer values.
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
func getFullDetails(usr User) {
	fmt.Println("The Name is from full details:", usr.name)
	fmt.Println("The Age is from full details:", usr.age)
}

////////////////////////////////////////////////////////////////////////////////
// Slice Demonstration
////////////////////////////////////////////////////////////////////////////////

// demonstrateSlices shows different types of slices and operations on them.
func demonstrateSlices() {
	fmt.Println("\n--- Slice Demonstration ---")

	// 1. Create a slice directly
	sliceArr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", sliceArr)
	fmt.Println("Length:", len(sliceArr)) // total number of elements
	fmt.Println("Capacity:", cap(sliceArr)) // total underlying array size

	// 2. Create a slice from another slice (sub-slicing)
	sliceArr1 := sliceArr[2:4] // includes index 2, excludes index 4
	fmt.Println("Sub-slice (sliceArr[2:4]):", sliceArr1)

	// 3. Append to a slice
	sliceArr = append(sliceArr, 6, 7)
	fmt.Println("After Append:", sliceArr)
	fmt.Println("Length After Append:", len(sliceArr))
	fmt.Println("Capacity After Append:", cap(sliceArr))

	// 4. Make a slice using make()
	// make([]Type, length, capacity)
	dynamicSlice := make([]int, 3, 5)
	dynamicSlice[0] = 10
	dynamicSlice[1] = 20
	dynamicSlice[2] = 30
	fmt.Println("Slice created using make():", dynamicSlice)
	fmt.Println("Length:", len(dynamicSlice))
	fmt.Println("Capacity:", cap(dynamicSlice))

	// 5. Copy a slice
	copySlice := make([]int, len(sliceArr))
	copy(copySlice, sliceArr)
	fmt.Println("Copied Slice:", copySlice)

	// 6. Nil slice example
	var nilSlice []int
	fmt.Println("Nil Slice:", nilSlice)
	fmt.Println("Is Nil?", nilSlice == nil)

	// 7. Slice of slices (multi-dimensional slice)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multi-dimensional Slice (Matrix):", matrix)
}

////////////////////////////////////////////////////////////////////////////////
// Main Function
////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println("Welcome To Golang")

	// Example of using a higher-order function
	higherOrderFunc(sumFunc)

	// Demonstrate different slice operations
	demonstrateSlices()

	// Anonymous function example
	func(name string) {
		fmt.Println("Hello", name)
	}("Sojib")

	// Printing user details
	getFullDetails(user1)

	// Demonstrating pointers
	usrDetails := getUserPointer(&arr)
	fmt.Println("The User Details are:", usrDetails)

	// Accessing struct methods
	fmt.Println("The name is:", user1.getName())

	// Fixed-size array
	fmt.Println(arr)

	// Run the condition checker
	Condition()
}
