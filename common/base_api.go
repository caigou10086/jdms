package common

import "github.com/gin-gonic/gin"

type BaseApi struct {
	Context *gin.Context
	Error   error
}

// MakeContext 设置上下文
func (b *BaseApi) MakeContext(c *gin.Context) *BaseApi {
	b.Context = c
	return b
}
