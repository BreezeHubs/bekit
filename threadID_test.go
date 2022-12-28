package main

import (
	"bekit/sys"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestThreadID(t *testing.T) {
	runtime.GOMAXPROCS(1) //设置线程数

	n := 10
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println("go1 threadID: ", sys.GetCurrentThreadIDByWin())
			for {
			}
		}()
	}

	go func() {
		count := 0
		for {
			time.Sleep(time.Second)
			count++
			fmt.Println("go2 threadID: ", sys.GetCurrentThreadIDByWin())
		}
	}()

	for {
	}
}

func TestCPUNum(t *testing.T) {
	sys.SetThreadNum(10)
	fmt.Println(sys.GetCPUNum(), sys.GetThreadNum())
}
