package main

import (
	"fmt"
	. "go/music/demo/mlib"
)

func Count (ch chan int) {
	fmt.Println("Count")
	// 写入
	ch <- 1
}

func main()  {
	mm := NewMusicManager()
	fmt.Println(mm.Len())


	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		fmt.Println("aa:", chs[i])
		// 每次打开一个新的goroutine，用一个形象的例子：本来一个人完成10个人交给的任务，现在是10个人完成10个人交给的任务
		go Count(chs[i])
	}

	for _, ch := range(chs) {
		// 读出
		<- ch
	}
}