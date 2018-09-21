package main

import (
	"fmt"
	//"time"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main()  {

	var ch = make( chan int, 1 )
	var arr []int

	wg.Add( 1 )
	go func( num int, at *[]int ){
		for i := 0; i < num; i++ {
			ch <- i
			*at = append( *at, <-ch )
		}
		wg.Done()
	}( 10, &arr )

	wg.Wait()
	fmt.Println( arr )
}