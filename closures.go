package main

import "fmt"

//闭包函数
func intSeq() func() int {
	i := 0

	return func() int {
		i +=1
		return i
	}
}

func main(){
	//返回一个匿名函数
	nextId := intSeq()

	fmt.Println(nextId())
	fmt.Println(nextId())
	fmt.Println(nextId())

	newInts := intSeq()

	fmt.Println(newInts())
}
