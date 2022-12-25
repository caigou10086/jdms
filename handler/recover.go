package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	// 被defer包裹的代码将会在return之后执行
	defer func() {
		// recover捕获异常
		if r := recover(); r != nil {
			log.Printf("panic:%v\n", r)
			// 打印异常信息
			debug.PrintStack()

			c.JSON(http.StatusOK, gin.H{
				"msg": r,
			})
			c.Abort()
		}
	}()
	c.Next()
}
