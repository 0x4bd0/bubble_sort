package main

import "fmt"

type Number interface{
	int | int16 | int32 | int64 | float32 | float64
}


func Sort[N Number](list []N) []N{
return list
}

func main(){
	
	 list1 := []int32{2,5,1,9,3,7,8,6,4,}
	 list2 := []float32{0.5,0.3,0.1,0.9,0.7,0.8,0.6,0.4,0.2, 1.0,}

	fmt.Println(Sort(list1))
	fmt.Println(Sort(list2))


}