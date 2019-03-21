package main

import (
	"fmt"
)

func main()  {
	//冒泡排序
	arr := []int{2,6,3,7,8}
	num := len( arr )

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++  {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println( arr )
}



