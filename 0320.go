package main

import (
	"fmt"
)

func main(){
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{6,7,8,9,10,1,1,1,1}

	//fmt.Printf( `%p`, s2 )

	copy( s2[2:4], s1[1:3] )	//copy 函数 把 s1 拷贝到s2当中去
	fmt.Println( s2 )


	//fmt.Printf( `%p`, s2 )	//还是会指向原有的底层数组


}

