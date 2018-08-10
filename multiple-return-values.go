package main

import "fmt"

//返回多个值 (int,int)表示返回两个值
func vals() (int,int) {
	return 3,7
}

func main(){
	a,b := vals()

	fmt.Println(a)
	fmt.Println(b)

	//如果只需要一个值可以是用空白定义符
	_,c := vals()
	fmt.Println(c)
}
