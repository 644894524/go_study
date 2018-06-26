package main

import (
	"fmt"
	"net/url"
	"io/ioutil"
	"strings"
	"net/http"
	"os"
)

type config struct {
	url string
	route string
}

var conf config = config{
	url : "http://10.10.20.99",
	route : "/admin/deploy/android/__auth",
}

//存取接受消息
var ch chan string = make( chan string )

func main(){
	names := [...]string{"sunlong@yixia.com","zhaona@yixia.com"}
	for key := range names {
		//异步执行
		go post_data( names[ key ] )
		fmt.Println( <-ch )
	}
}

func post_data( name string ){
	host := conf.url + conf.route
	v := url.Values{}
	v.Set( "name", name )
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := http.Client{}
	req, err := http.NewRequest( "POST", host, body )
	if err != nil {
		os.Exit( 500 )
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//开始请求发送数据
	resp, err := client.Do( req )
	if err != nil {
		os.Exit( 500 )
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	rescode:= string(data)
	ch <- rescode
}