package main

import (
	"fmt"
)

func func1() {
	i := 1
Here:
	fmt.Printf("%d", i)
	i++
	if i > 10 {
		return
	}

	goto Here
}

func sum_total(int limit) {
	sum := 0
	for i := 0; i < limit; i++ {
		sum = sum + i
	}

	return sum

}
func main() {
	// if err := Chmod(0644); err != nil {
	// 	fmt.Printf(err)
	// }

	if true && true {
		fmt.Printf("true")
	}

	if !false {

		fmt.Printf("true")
	}
	func1()
	fmt.Printf("total is %d", sum_total(50))
}
