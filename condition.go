package main

import "fmt"




 func Condition(){
	var age int;
	fmt.Println("Please Input Your Age:");
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("Oops, You can not Do this")
	}else {
	fmt.Println("Oops, You can Do this")
	}
 }

