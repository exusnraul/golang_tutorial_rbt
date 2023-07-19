package main

import "fmt"

func main() {

	func1("Rahul", 55892)

}

func func1(Name string, id int) {
	fmt.Println(Name, id)
	return Name
}

func func2(Name string) int {
	return 100

}

//func returning multiple values

func func3(id int) (string, int) {
	var Name string = "raul"
	var salary = 100000
	return Name, salary
}
