package comm

import (
	"github.com/gin-gonic/gin"
)

// Cors 中间件
func Cors(ctx gin.Context) {
	ctx.Header("", "")
}
