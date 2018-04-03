package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main(){
	//读取所有文件
	data, _ := ioutil.ReadFile( "./map.go" )
	fmt.Println( string( data ) )

	//只读取文件的一部分
	f,err := os.Open( "./map.go" )	//open ./map1.go: no such file or directory 会报错
	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}

	content := make( []byte, 10 )
	n1,err := f.Read( content )
	fmt.Printf( "%d bytes : %s \n", n1, string( content ) )
	f.Close()

	//读取行

}
