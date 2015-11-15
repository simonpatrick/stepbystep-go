package main
import (
	"fmt"
	"math"
)

type OOPSample struct {
	a,b,c int
}

func (oopSample OOPSample) sum() int{
	return  oopSample.a+oopSample.b+oopSample.c
}

func (oopSample *OOPSample) sumAll() int{
	return oopSample.a+oopSample.b+oopSample.c
}

type MyFloat float64 //enhance type anytime
func (f MyFloat) abs() float64{
	if(f<0){
		return float64(-f)
	}else {
		return float64(f)
	}
}

func main() {
	fmt.Println(OOPSample{1,2,3}.sum())
	fmt.Println(MyFloat((-math.Sqrt2)).abs())
}