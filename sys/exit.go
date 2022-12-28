package sys

import (
	"os"
	"os/signal"
	"syscall"
)

type eSignal struct {
	c      chan os.Signal
	isExit bool
}

func NewListenExitSignal() *eSignal {
	// 从这里开始优雅退出监听系统信号，强制退出以及超时强制退出。
	c := make(chan os.Signal, 1)

	e := &eSignal{c: c}
	go e.wait()
	return e
}

func (e *eSignal) wait() {
	//windows
	signal.Notify(e.c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGTERM)

	//linux & mac
	//signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
	//	syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL,
	//	syscall.SIGTRAP, syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM)

	select {
	case <-e.c:
		go func() {
			select {
			case <-e.c:
				os.Exit(1) //再次监听退出信号
			}
		}()
		e.isExit = true
	}
}

func (e *eSignal) IsExit() bool {
	return e.isExit
}
