package main

import "os" //parse commandline
import "fmt"
import "simplemath"
import "strconv"

var Usage =func ()  {
  fmt.Println("USAGE: calculation sample")
  fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
  args:=os.Args[1:]
}
