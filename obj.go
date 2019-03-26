package main

import "fmt"

type person struct {
	code string
	name string
	age int
	city string
}

func ( person *person ) say () string {
	var response string = fmt.Sprintf( "my name is %s, age : %d, city : %s", person.name, person.age, person.city )
	return response
}

func main(){
	var stu = person{"220322", "sunlong", 28, "jilin"}
	var res string = stu.say()
	fmt.Println( res )

	/*
	var str string = "hello world"
	zhizhen := &str
	*zhizhen = "hello"
	fmt.Println( *zhizhen, str )
	*/


	var str string = "hello world"
	var strone string = str
	strone = "hello"
	fmt.Println( strone, str )


	slice2 := make([]int, 3, 5)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	slice2 = append( slice2, 4 )
	slice2 = append( slice2, 5 )
	slice2 = append( slice2, 6 )
	fmt.Println( len( slice2 ), cap( slice2 ) )


	slice3 := []int{1,2,3}
	slice3 = append( slice3, 4 )
	fmt.Println( slice3 )
}