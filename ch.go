package main

import (
	"time"
	"fmt"
)

func print( ch chan string ){
	strone := "hello world"
	time.Sleep( 10000 * time.Millisecond )		//暂停10S
	ch <- strone
}

func main(){
	fmt.Println( "start" )
	ch := make( chan string )
	go print( ch )
	strtwo := <- ch
	fmt.Println( strtwo )
	fmt.Println( "end" )
}