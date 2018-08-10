package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main(){
	//创建新的结构体
	fmt.Println(person{"Bob",20})

	fmt.Println(person{name:"Alice",age:30})

	//省略的字段被初始化为零值
	fmt.Println(person{name:"Fred"})

	//&前缀生成一个结构体指针
	fmt.Println(&person{name:"Ann",age:40})

	s := person{name :"	Sean",age:50}
	//使用点来访问结构体的字段
	fmt.Println(s.name)

	//也可以对结构体指针使用.-指针会被自动解引用
	sp := &s
	fmt.Println(sp.age)

	//结构体是可变的
	sp.age = 51
	fmt.Println(sp.age)

}
