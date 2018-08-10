package main

import "fmt"

func main(){
	//创建关联数组(map)的方法
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:",m)

	v1 := m["k1"]
	fmt.Println("v1:",v1)
	//map的长度
	fmt.Println("len:",len(m))

	//从map中删除一个值
	delete(m,"k2")
	fmt.Println("map:",m)

	//_的作用是如果k2不存在测prs为false,否则是0
	_,prs := m["k2"]
	fmt.Println("prs:",prs)

	//map的初始化并赋值
	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)

}
