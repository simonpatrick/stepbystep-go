package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func reverse(a, b int) (int, int) {
	return b, a
}


func main() {

	fmt.Println(sum(2, 3))
	fmt.Println(reverse(2, 3))
}
