package main

//切片练习
import "fmt"


func main(){
	slice1 := [5]int{ 1, 3, 5, 7, 9 }
	fmt.Print( len(slice1), cap(slice1) )



	slice2 := slice1[3:]
	if slice2 == nil {	//切片只能和nil进行对比
		fmt.Print( `呵呵` )
	}

	slice2[0] = 8	//创建一个新的切片执行同样的底层数组 , 包含了指向数组的指针
	slice2 = []int{10,20}
	fmt.Print( slice1, slice2 )



	slice3 := make( []int, len(slice1) * 2 )	//翻倍扩容
	//slice3 = slice1		//cannot use slice1 (type [5]int) as type []int in assignment	不能够直接进行赋值

	for i := range slice1 {
		slice3[i] = slice1[i]
	}

	fmt.Print( "\r\n", slice1, slice3 )
	fmt.Print( "\r\n", cap(slice1), cap(slice3) )


	slice3[5] = 2
	fmt.Print( slice3 )



	//slice2 := make( []int, 5 )
	//fmt.Print( slice1[:] )
	//copy( slice1[:], slice2 )
	//fmt.Print( slice2 )


}
