package simplemath

import "math"

func Sqrt(i int)  {
  v:=math.Sqrt(float64(i))
  return int(v)
}
