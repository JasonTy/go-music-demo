package main

import (
	"fmt"
	. "go/music/demo/mlib"
)

func main()  {
	mm := NewMusicManager()
	fmt.Println(mm.Len())
}