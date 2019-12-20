package http_server

import (
	"daily/reverse_tunnel/server/cfg"
	"daily/reverse_tunnel/server/http_server/handler"
	"daily/reverse_tunnel/server/socket"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.StaticFile("/client/private_key/mips", cfg.Config().SSHConfig.ReverseKeys.MIPSPrivateKeyPath)
	router.StaticFile("/client/private_key/normal", cfg.Config().SSHConfig.ReverseKeys.XPrivateKeyPath)
	router.StaticFile("/server/public_key", cfg.Config().SSHConfig.ForwardKeys.PublicKeyPath)

	router.GET("/health", handler.HealthHandler)

	router.GET("/ws", socket.SocketHandler)

	router.GET("client_private_key", handler.ClientPrivateKeyHandler)

	router.GET("/members", handler.GetMembersHandler)

	// 新建一个端口映射
	router.PUT("/port_mapping", handler.CreateTunnelHandler)
	// 删除一个端口映射
	router.DELETE("/port_mapping", handler.CreateTunnelHandler)

	router.POST("/progress", handler.ProgressHandler)

	if err := router.Run("0.0.0.0:6998"); err != nil {
		panic(err.Error())
	}

}
