package main

import (
	"fmt"
)

func main() {

	s := "hello world"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	//multiple line input
	s3 := "Line1" +
		"Line2"
	fmt.Printf("%s\n", s3)

	s4 := "Line1"
	+"Line2" //error here
	fmt.Printf("%s\n", s4)
}
