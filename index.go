
package main

import "fmt"

func main() {
	result := total("Huu Hung")
	fmt.Println(result)
	
	w,h,area,name := rectInfo(100,200)
	fmt.Println("width=", w)
	fmt.Println("height=", h)
	fmt.Println("area=", area)
	fmt.Println("name =", name)
	

	w,h,isSquare := rectInfo2(200,200) 
	if isSquare {
		fmt.Printf("This is square")
	} else {
		fmt.Println("width",w)
		fmt.Println("height", h)
	}
}

func hi() {
	fmt.Println("毎日ITを勉強してる頑張りましょう")
}

func sayHello(name string) {
	fmt.Println("こんにちは",name)
}

func total( name string ) string {
	result := fmt.Sprintf("SayHello %s", name)
	return result
}

// Multiple return values
func rectInfo(w, h int) (int, int, int, string) {
	area := w * h
	name := "ITを勉強してる"
	return w, h, area, name
}
	
// Name return values
func rectInfo2(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w,h,isSquare
}

