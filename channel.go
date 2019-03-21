package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main()  {
	wg.Add( 1 )			//有多少个子进程
	go func( a int ){
		for i:=0; i<a; i++{
			fmt.Println( i )
		}
		wg.Done()				//进程结束
	}( 10 )

	wg.Wait()
	fmt.Println( "进程结束" )
}