package main

import (
	"fmt"
	//"time"
	"sync"
	//"os"
)

var (
	wg sync.WaitGroup
)

func main()  {

	var ch = make( chan int, 1 )
	var arr []int
	var num int = 10

	wg.Add( 1 )
	go func( num int, at *[]int ){
		for i := 0; i < num; i++ {
			ch <- i
			*at = append( *at, <-ch )
		}
		wg.Done()
	}( num, &arr )

	go func( con *[]int ) {
		for {
			fmt.Println( len( *con ) )
		}
	}( &arr )

	wg.Wait()

	fmt.Println( arr )
}



