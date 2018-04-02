package main

import (
	"os"
	"fmt"
)

func main(){
	file := `./map1.go`
	f,err := os.Open( file )	//open ./map1.go: no such file or directory 会报错
	check( err )
	content := make( []byte, 30 )
	n1,err := f.Read( content )
	fmt.Println( string( n1 ) )
	f.Close()
}


func check( e error ){
	if e != nil {
		panic( e )
	}
}