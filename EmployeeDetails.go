package main

import(
	"fmt"
)

func main(){

	fmt.Println("---------- Welcome to Zopdev ------------\n")

	var name string 
	var age int
	var cgpa float32

	fmt.Println("Input of User details: Name , Age , Cgpa")
	fmt.Scanf("%s %d %f",&name,&age,&cgpa)

	fmt.Printf("Click 1 to print Employee Details and 2 for exit : ")
    var t int ;
	fmt.Scanf("%d",&t)

	if(t==1){
       	fmt.Println("\n------- Employee Details --------")
		fmt.Println("Name :",name , "\nAge  :",age, "\nCgpa :",cgpa,"\n")
        
	} else{
        
	}

	    fmt.Println("----- Thank You ----")
	



}