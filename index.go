package main

import "fmt"

func main() {
	// Cách sử dụng if else 
	number := 10 

	// if condition {//code here}
	if number == 10 {
		fmt.Println("number = 10")
	}

	// if else
	if number < 10 {
		fmt.Println("number < 10")
	}else {
		fmt.Println("number = 10")
	}

	// if statement, condition {//code}
	if a:= 100; a==100 {
		fmt.Println("a = 100")
	}

	// switch case
	number2 := 20
	switch number2 {
	case 1,200,300:
		fmt.Println("number = 1")
	case 2:
		fmt.Println("number = 2")
	case 3:
		fmt.Println("number = 3")
	case 20:
		fmt.Println("number = 4")
	default:
		fmt.Println("unknown")

	}

	number3 := 30
	switch {
	case number3 > 10:
			fmt.Println("number > 10")
	case number3 == 10:
			fmt.Println("number == 10")
	}

	// FallThrough Tiếp tục Math
	number4 := 10
	switch number4 {
		case 1:
			fmt.Println("number = 1")
			fallthrough
		case 10: 
			if number == 10 {
				goto handleNumberEqual10
			}	
			handleNumberEqual10:
				fmt.Println("handle for case = 10")
				break
		case 2:
			fmt.Println("number = 1")
			fallthrough
		case 3:
			fmt.Println("number = 3")
			fallthrough
		case 4:
			fmt.Println("number = 4")
		
	}

}