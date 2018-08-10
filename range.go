package main

import "fmt"

//range 相当于php中的foreach
func main(){
	nums := []int{1,2,3}
	sum := 0
	for _,num := range nums {
		fmt.Println("num:",num)
		sum += num
	}

	fmt.Println("sum:",sum)

	for i,num := range nums {
		if num == 3 {
			fmt.Println("index:",i,"value:",num)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}

	//也可以循环字符串查询字符串对应的unicode值
	for i,c := range "go" {
		fmt.Println(i,c)
	}
}
