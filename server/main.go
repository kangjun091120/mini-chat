package main

import (
	"chatter-server/mysql"
	"chatter-server/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	go ws.WebsocketManager.Start()
	mysql.InitDB()

	router := gin.Default()

	// server 核心
	server := router.Group("/server")

	type User struct {
		Name     string `json:"name" form:"name"`
		Password string `json:"password" form:"password"`
	}
	// 注册部分
	server.POST("/register", func(context *gin.Context) {
		var user User
		_ = context.BindJSON(&user)
		if !mysql.Checkuser(user.Name) {
			mysql.Register(user.Name, user.Password)
			context.JSON(http.StatusOK, gin.H{
				"code":   200,
				"status": "SUCCESS",
				"error":  "",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":   412,
				"status": "FAIL",
				"error":  "name has already be used",
			})
		}
	})

	//登陆部分
	server.POST("/login", func(context *gin.Context) {
		var user User
		_ = context.BindJSON(&user)
		if !(mysql.Login(user.Name, user.Password) == 0) {
			context.JSON(http.StatusOK, gin.H{
				"code":   200,
				"status": "SUCCESS",
				"error":  "",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":   404,
				"status": "FAIL",
				"error":  "name or password not found",
			})
		}
	})

	router.GET("/", func(c *gin.Context) {
		ws.WebsocketManager.SendGroup("test", []byte("aaa"))
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// websocket 升级协议
	router.GET("/ws/:user", ws.WebsocketManager.WsClient)
	_ = router.Run("0.0.0.0:81")
}
