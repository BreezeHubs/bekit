package middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var errStr string
				if stack {
					//打印错误堆栈信息
					log.Printf("panic: %v\n", r)
					debug.PrintStack()

					errStr = "服务器内部错误" + fmt.Sprintf("panic: %v\n", r)
				} else {
					errStr = "服务器内部错误"
				}

				//封装通用json返回
				c.JSON(500, gin.H{500, errStr, nil})
			}
		}()
		c.Next()
	}
}
