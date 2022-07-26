package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

func main() {
	go func() {
		println("Hello, world. 1")
	}()

	println("Hello, world. main")

	go running()

	var input string
	fmt.Scanln(&input)
}
