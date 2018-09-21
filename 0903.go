package main

import "fmt"

type readOp struct {
	key string
	val string
}

func ( i readOp ) Str() string {
	str := fmt.Sprintf( "%s %s", i.key, i.val )
	return str
}

func main()  {
	obj := make( chan readOp )
	obj <- readOp{ "haha", "haahj" }
	close( obj )
}