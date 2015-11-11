package main

import "fmt"

func remove(arr []int,i int)  {
  return append(arr[:i],arr[i+1:])
}

func testRemove()  {
  a:=[]int{1,2,3,4,5}
  a=remove(a,0)
  fmt.Println(a)

  a=remove(a,3)
  fmt.Println(a)
  a=remove(a,1)
  fmt.Println(a)
}

func main() {
  testRemove()
}
