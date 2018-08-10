package main

import "fmt"

//函数的定义
//func 函数名(参数1 类型,参数2 类型) 返回类型 {逻辑代码}
func plus(a int, b int) int {
	return a + b
}

func main(){
	res := plus(1,2)

	fmt.Println("1+2=",res)
}
