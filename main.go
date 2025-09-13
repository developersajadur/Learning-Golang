package main

import "fmt"


// init function is a special function in golang
// it will execute before the main function
// we can use it to initialize variables or setup configurations
// we can have multiple init functions in a single file or package
// they will execute in the order they are defined
  func init(){
	fmt.Println("The Init This should be the first to execute")
  }

func main()  {
	fmt.Println("Welcome To golang")
	higherOrderFunc(sumFunc)
	 func (name string){
	fmt.Println("Hello", name)
 }("Sojib")

}

// higherOrderFunc is a function that takes another function as an argument
func higherOrderFunc(HOF func(int, int)) {
	HOF(10,20)
}

 func sumFunc(num1 int, num2 int){
	fmt.Println("The Sum Is:", num1 + num2)
 }



