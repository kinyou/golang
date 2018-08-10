package main

import "fmt"

//go的变量
func main(){
	var a string = "initial"
	fmt.Println(a)

	//一次性申明多个变量
	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	//整型的变量申明不赋值默认是0
	var e int
	fmt.Println(e)

	//go会自动推导变量类型,申明并赋值的申明语法
	f := "short"
	fmt.Println(f)

	//字符串变量申明不赋值默认是""
	var g string
	fmt.Println(g == "")

	//bool的变量申明不赋值默认是false
	var bool bool
	fmt.Println(bool)


}