package main

import (
	"bekit/sys"
	"fmt"
	"time"
)

func main() {
	s := sys.NewListenExitSignal()

	for {
		if s.IsExit() {
			break
		}
		fmt.Println(s.IsExit())
		time.Sleep(2 * time.Second)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(s.IsExit())
}
