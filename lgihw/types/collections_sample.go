package main

import "fmt"


func main() {
	a :=[]int{1,2,3,4,5}
	fmt.Println(a[2:3])
	fmt.Println(a[:3])
	fmt.Println(a[2:])
}