package main

import "fmt"

func main() {
	fmt.Println("Message from branch main ...")
	fmt.Println("Message from feature development...")
	println("1 + 5 = ", sumtwonumbers(1, 5))
	println("2 * 5 = ", subtractiontwonumber(2, 5))
}

func sumtwonumbers(x int, y int) int {
	return x + y
}

func subtractiontwonumber(x int, y int) int {
	return x * y
}
