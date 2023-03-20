package main

import (
	"fmt"
	"math"
)

func main() {
	// Boolean
	myBoll := true
	fmt.Println(myBoll)

	mySecond := false
	fmt.Println(mySecond)
	
	// String
	myString := "毎日ITを勉強してる"
	fmt.Println(myString)
	
	//int kiểu số nguyên
	myInt := 123
	fmt.Println(myInt)
	
	//int 8,16,32,64 
	// 1.Range
	// 2. Bits
	
	//1. Range int8 (-128 - 127)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	//1. Range int16 (-32768 - 32768)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	//1. Range int32 (-2147483648 - 2147483647)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	//1. Range int64 (-9223372036854775808 - 9223372036854775807)
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	// // 2 Bits 
	// fmt.Println("===========")
	// fmt.Println(bits.OnesCount8(math.MaxUint8))
	// fmt.Println(bits.OnesCount16(math.MaxUint16))
	// fmt.Println(bits.OnesCount32(math.MaxUint32))
	// fmt.Println(bits.OnesCount64(math.MaxUint64))

	// =================================================
	// uint	Số nguyên dương

	var myInt2 uint = 10  // uint là gán số nguyên dương nếu số này âm thì báo lỗi)
	fmt.Println(myInt2)
	
	// =================================================
	// Byte: bản chất của byte là unit8
	var myByte byte = 255
	fmt.Println(myByte)
	// fmt.Println(math.MaxUint8)
	fmt.Printf("T", myByte)
	
	// =================================================
	var a byte = 'E'
	fmt.Printf("%X", a)

	// ===================
	fmt.Println("================")
	// float Kiểu dữ liệu số thực
	var myFloat float64 = 10.01
	fmt.Println(myFloat)

	// Complex Kiểu dữ liệu số phức z = a + bi (a là phần nguyên b là phần ảo)
	var z complex64 = 10 + 1i
	fmt.Println(z)

	var z2 complex64 = 10 + 2i
	fmt.Println(z + z2)
	
	// ===========================================================
	// Rune
	var myString2 string = "Nhẫn"

	runes := []rune(myString2)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}	 

	var myRune rune = 'A'
	fmt.Printf("%T", myRune)

	
	var c int = 1
	var d float64 = 2.1
	fmt.Println(float64(c) + d)

	const PI = 3.14159
	fmt.Println(PI)
} 

