package handler

import (
	"daily/reverse_tunnel/server/cfg"
	"github.com/gin-gonic/gin"
)

func ClientPrivateKeyHandler(c *gin.Context) {
	c.File(cfg.Config().SSHConfig.ReverseKeys.XPrivateKeyPath)
}
