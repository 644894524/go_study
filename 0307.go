//定义了包名,每个项目都必须包含main
package main

//导入包
import(
	"fmt"
	//"os/exec"
)

const host_name string = `127.0.0.1`

func main(){

	//常量声明


	//三种声明方式
	var host string = `127.0.0.1`
	var pwd = `1q2w3e4r`
	user := `root`
	var port string = `3306`

	host = `1q2w3e4r`

	//Sprintf 不会进行标准输出
	var command string = fmt.Sprintf( "/usr/local/mysql/bin/mysql -h %s -P %s -u %s -p'%s'", host, port, user, pwd )
	fmt.Print( command + "\r\n" )

	//cmd := exec.Command( command )
	//err := cmd.Run()

	//返回报错了
	//fmt.Print( err )

	fmt.Print( host_name )
}

