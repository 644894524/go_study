package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"os"
)

func main(){
	//get 请求
	url := "http://sl.test.com/api.php?name=sunlong"
	client := &http.Client{ Timeout:1 * time.Second }
	res, err := client.Get( url )
	check( err )
	res.Header.Set( "Content-Type", "application/x-www-form-urlencoded" )	//设置header

	result, _ := ioutil.ReadAll( res.Body )	//返回beta
	res.Body.Close()

	//fmt.Printf( `%s`, result )
	msg := string( result )
	fmt.Println( msg )
}

func check( err error ){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println( r )
			os.Exit( 1 )
		}
	}();

	if err != nil {
		panic( `request timeout` )
	}
}
