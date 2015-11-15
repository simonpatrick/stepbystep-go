package main
import "fmt"

func main() {
	pow :=[]int{1,2,4,8,16,32,64,128}
	for i,v:= range pow{
		fmt.Printf("2**%d=%d \n",i,v)
	}

	a := map[string]string{};
	a["test"]="test"
	fmt.Println(a)
}
