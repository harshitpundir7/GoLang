package main

import (
	"fmt"
	"time"
)

func main (){
	i :=5
	switch i {
	case 1 : 
		fmt.Println("One")
	case 2 :
		fmt.Println("two")
	case 3: 
		fmt.Println("Three")
	default: 
		fmt.Println("Other")
	}

	//multiple condition Switch
 fmt.Println(time.Now().Weekday())
	 switch time.Now().Weekday(){
	 case time.Saturday,  time.Sunday: 
		fmt.Println("It's Weekend")
	 default : 
	  fmt.Println("Its workDay")
	}

	//switch case through function
	whoAmI := func(i interface{}){
		switch t := i.(type) {
		case int :
			fmt.Println("its a Integer")
		case string : 
			fmt.Println("its a String")
		case bool : 
			fmt.Println("its a bool", t)
		default :
			fmt.Println("other")
		}
	}
	whoAmI(43)
	whoAmI("Harshit")
	whoAmI(true)
	whoAmI(7.3536)
}