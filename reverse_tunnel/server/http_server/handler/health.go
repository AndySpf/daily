package handler

import "github.com/gin-gonic/gin"

// HealthHandler 客户端心跳，检测服务端ssh反向代理是否存活。客户端需根据当前状态选择是否重新建立连接
func HealthHandler(c *gin.Context) {
	c.String(200, "ok")
}
