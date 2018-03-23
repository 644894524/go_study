package main

import (
	"fmt"
	//"os"
)

func main()  {

	/*s1 := []string{ `a`, `b`, `c`, `d`, `e` }
	fmt.Printf( `\n %p => %v`, s1, s1 )

	//这里指向的是同一个内存地址
	//s2 := s1
	//fmt.Printf( `%p => %v`, s2, s2 )
	//fmt.Print( "\n" )


	sa := s1[2:5]
	//s1 = append(s1, `f` )
	sa[0] = `g`


	fmt.Print( s1, sa, "\n" )	//s1发生了更改
	//os.Exit( 0 )

	//切片
	s2 := make( []int, 3, 6 )
	fmt.Println( s2 )
	s2[0] = 1
	fmt.Printf( "\n %p => %v", s2, s2 )

	s2 = append(s2, 4, 5, 6)
	fmt.Printf( "\n %p => %v", s2, s2 )	//内存地址不会发生改变

	s2 = append(s2, 9)
	fmt.Printf( "\n %p => %v", s2, s2 )	//内存地址发生改变*/


	s3 := []int{1, 2, 3, 4, 5}
	i1 := s3[2:5]
	i2 := s3[1:3]
	fmt.Print( "\n", i1, i2 )
	i1 = append(i1, 8)
	i1[0] = 9
	fmt.Print( "\n", i1, i2 )	//i1 和 i2 共同发生改变  , 40行注释放开则不会改变*/
}
