package handler

import (
	"daily/reverse_tunnel/server/socket"
	"github.com/gin-gonic/gin"
)

func GetMembersHandler(c *gin.Context) {
	res := socket.GetOnlineMembers()

	c.JSON(200, res)
}
