package main

import "fmt"

func main() {
	weekdays := []string{"Mon", "Tue", "wed", "Thu", "Fri"}

	weekdays1 := append(weekdays[:2], weekdays[3:]...)

	fmt.Println(weekdays)
	fmt.Println(weekdays1)
}
