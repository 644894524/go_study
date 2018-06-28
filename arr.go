package main

import "fmt"

func main(){

	var arr = [...]int{1,2,3,4,5,6,10,11}
	fmt.Println( arr )


	var arr1 = [3]int{}
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println( arr1 )



	//边长数组
	arr3 := make( []int, 8, 10)
	var len int = 8
	for i := 0; i < len; i++ {
		arr3 = append( arr3, i )
	}

	fmt.Println( arr3 )



	arr4 := []int{}
	var len4 int = 8
	for i := 0; i < len4; i++ {
		arr4 = append( arr4, i )
	}

	fmt.Println( arr4 )
}

