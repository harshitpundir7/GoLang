package main

import "fmt"

func main(){
	age := 16
	if age>=18 {
		fmt.Println("Person is Adult")
	} else if age>=12{
		fmt.Println(("Persin is Teenager"))
	} else {
		fmt.Println(("person is kid"))
	}

	var role =  "Admin"
	var hasPermission = true
	if role =="Admin" || hasPermission {
		fmt.Println("yes")
	}


	// we can declare an veriable inside the iF-Else loop
	if age:=15; age>=18 {
		fmt.Println("Person is Adult ",  age)
	} else if age<=12 {
		fmt.Println("Person is Teenager", age)
	}

	// go does'nt have ternary operator, We have to use Normal if Else. 
}