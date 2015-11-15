package main
import "fmt"

func main() {
	my_channel :=make(chan int)
	// put 5 to channel
	my_channel <-5
	var my_recieved_value int =<- my_channel
	fmt.Println(my_recieved_value)
}
