package main

import (
	"encoding/json"
	"os"
	"fmt"
)

func main() {

	//简单的
	jsonData1 := []byte( `["a","b","c","d","e"]` )
	var arr []string
	err := json.Unmarshal( jsonData1, &arr )
	if err != nil {
		os.Exit( 20 )
	}
	fmt.Println( arr[0] )


	jsonData2 := []byte( `{"code":"200","msg":"this is msg"}` )		//注意类型  eg  code : 200 变成整形了，则不能解析
	var map1 map[string]string
	json.Unmarshal( jsonData2, &map1 )
	fmt.Println( map1 )


	//必须是大写的 , 最好是变成map形式
	jsonData3 := []byte( `{"Code":"200","Msg":"this is msg"}` )
	type conf struct {
		Code int
		Msg string
	}
	var jsonformat conf
	json.Unmarshal( jsonData3, &jsonformat )
	//fmt.Println( jsonformat[0] )


	//第三方包

}