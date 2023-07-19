package main

import "fmt"

func main() {

	var exp = outerfunc()

	fmt.Println(exp())
}

func outerfunc() func() int {
	exp := func() int {
		return 1000
	}
	return exp
}
