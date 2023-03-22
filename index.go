package main

import "fmt"

func main() {
	// Khai báo array
	var myArray [4]int
	fmt.Println(myArray)

	// Gán giá trị cho array
	myArray[0] = 123
	myArray[1] = 456
	myArray[2] = 789
	myArray[3] = 10
	fmt.Println(myArray)

	// Khai báo 1 array có khỏi tạo giá trị

	array := [...]int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	fmt.Println(array)

	array1 := [4]int{100}
	fmt.Println(array1)

	// size mảng
	fmt.Println(len(array))

	// Khai báo mảng không cần set size
	array3 := [...]string{"Javascript", "PHP", "Ruby", "React", "Node", "1"}
	fmt.Println(array3[2])

	// array là value type không phải là reference type

	countries := [...]string{"VN", "US", "FRANCE"}
	copyCountries := countries

	copyCountries[0] = "CANADA"

	fmt.Println(countries)
	fmt.Println(copyCountries)

	fmt.Println("===============")
	// loop array
	// Cách 1
	for i := 0; i < len(copyCountries); i++ {
		fmt.Println(copyCountries[i])
	}

	fmt.Println("===============")
	// Cách 2
	for index, value := range countries {
		fmt.Printf("index=%d value=%s", index, value)
		fmt.Println()
	}

	fmt.Println("===============")

	// Khi chỉ muốn làm việc với value không muốn làm việc với index
	// blank identifier
	for _, value := range countries {
		fmt.Printf("value = %s", value)
		fmt.Println()
	}

	// mảng hai chiều,  [hàng][cột] [row][col]

	matrix := [4][2]int{ //4 hàng hai cột
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(matrix)

	fmt.Println("===============")

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], "  ")
		}
		fmt.Println()
	}

}
