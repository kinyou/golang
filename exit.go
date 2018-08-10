package main

import (
	"fmt"
	"os"
)

func main(){

	defer fmt.Println("使用os包里面的defer修饰的语句永远不会执行")

	os.Exit(3)
}
