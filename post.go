package main

import (
	"net/http"
	"time"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main(){
	var host string = "http://sl.test.com/api.php"
	postV := "name=lisi"
	client := &http.Client{ Timeout:1 * time.Second }
	res, err := client.Post( host, "application/x-www-form-urlencoded", strings.NewReader( postV ) )
	check( err )

	result, _ := ioutil.ReadAll( res.Body )
	res.Body.Close()
	msg := string( result )
	fmt.Println( msg )
}

func check( err error ){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println( r )
			os.Exit( 1 )
		}
	}()
	if err != nil {
		panic(`Request timeout`)
	}
}
