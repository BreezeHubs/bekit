package sys

import (
	"runtime"
	"syscall"
)

func GetCurrentThreadIDByWin() int {
	var (
		user32   *syscall.DLL
		threadID *syscall.Proc
		err      error
	)

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		return -1
	}

	threadID, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		return -1
	}

	pid, _, _ := threadID.Call()
	return int(pid)
}

func SetThreadNum(n int) int {
	return runtime.GOMAXPROCS(n)
}

func GetThreadNum() int {
	return runtime.GOMAXPROCS(0)
}

// GetCPUNum 当前系统的CPU核数量
func GetCPUNum() int {
	return runtime.NumCPU()
}
