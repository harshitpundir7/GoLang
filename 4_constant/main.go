package main

import "fmt"

const age = 20

func main(){
	const name string = "Harshit"
	//name = "JavaScript" // we can not change this like this because its constant.
	fmt.Println(name)
	fmt.Println(age)
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}