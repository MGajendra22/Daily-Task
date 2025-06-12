package main

import (
	"fmt"
)

func main(){

	var i int =20
	var j int =30

	ptr:=&i

	defer fmt.Println(*ptr)

	ptr=&j

	fmt.Println("Value of ptr :",*ptr)

	switch {

	case i==20:
		fmt.Println("Case 1 executed")
	case i==20:
		fmt.Println("Case 2 executed")	
	default:
		fmt.Println("Exit")	
	}

}