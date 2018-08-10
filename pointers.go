package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

//指针
func zeroptr(iptr *int) {
	*iptr = 0
}

func main(){
	i := 1
	fmt.Println("initial:",i)

	//这个不可以改变,只是把参数copy一份
	zeroval(i)
	fmt.Println("zeroval:",i)

	//指针可以改变变量的值
	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	fmt.Println("pointer:",&i)
}
