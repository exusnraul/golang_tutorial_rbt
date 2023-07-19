package main

import "fmt"

func main() {
	var map1 = map[string]int{"java": 10, "Node": 20, "Go": 50}
	fmt.Println(map1["Go"])
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
