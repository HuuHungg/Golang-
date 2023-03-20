package main

import "fmt"

func main() {
	// loops 
	// for init, condition, post

	// break, continue
	for i := 0; i < 10; i++ {
		if(i == 4) {
			continue
		}
		fmt.Println(i)
	} 
	fmt.Println("out of loop")
	
	// While trong Golag
	j := 0
	for ; j < 5; {
		fmt.Println(j)
		j++
	}
	// for init, condition
	for i,j := 1, 2; i < 10 && j < 10; i,j = i + 1, j + 1 {
		fmt.Println(i)
		fmt.Println(j)
	}

}