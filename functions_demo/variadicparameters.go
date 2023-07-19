package main

import "fmt"

func main() {
	calculatemarks("asdf")
	calculatemarks("asdf", 80, 23, 23)
}

func calculatemarks(Name string, marks ...int) {
	fmt.Println(Name)
	fmt.Println(marks)
	fmt.Println("++++++")
}
