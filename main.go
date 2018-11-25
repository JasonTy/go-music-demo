package main

import (
	"fmt"
	. "go/music/demo/mlib"
	"sync"
)
var once sync.Once
var l sync.Mutex
var a = 0
func PrintStr()  {
	defer l.Unlock()
	fmt.Println("Count:", a)
	a++
}

func Count (ch chan int) {
	l.Lock()
	PrintStr()
	ch <- 1
}

func main()  {
	mm := NewMusicManager()
	fmt.Println(mm.Len())


	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		// 每次打开一个新的goroutine，用一个形象的例子：本来一个人完成10个人交给的任务，现在是10个人完成10个人交给的任务
		go Count(chs[i])
	}

	for _, ch := range(chs) {
		// 读出
		<- ch

	}
}