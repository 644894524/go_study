package main

import (
	"fmt"
	"sync"
)

func main (){
	var count int = 100
	sy := sync.WaitGroup{}
	for i := 0; i <= count; i ++ {
		sy.Add(1)
		go func( i int ) {
			fmt.Println( i )
			sy.Done()
		}( i )
	}

	sy.Wait()
	fmt.Println( "end" )
}