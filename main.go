package main

import "fmt"

func main()  {
	fmt.Println("Welcome To golang")
	callTheFunc := helloFunc(2, 3)
	fmt.Println(callTheFunc)

	callGetValueAndSum := getValueAndSum()
	fmt.Println("Total Of Sum After The Inputs Is:", callGetValueAndSum)
}
