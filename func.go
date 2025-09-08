package main
import "fmt"

func helloFunc(num1 int, num2 int) int {
	return num1 + num2
}



func getValueAndSum() int{
	var num1 int = 10
	var num2 int = 20

	fmt.Println("Input The Value Of Num1")
	// fmt.Scanln(&num1)
	fmt.Println("Input The Value Of Num2")
	// fmt.Scanln(&num2)

	return num1 + num2
}