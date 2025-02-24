package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic for Loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//continue and break Statement
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//Range in Go lang
	for i := range 3 {
		fmt.Println(i)
	}

	//infinte loop
	for {
		fmt.Println("1")
	}
}
