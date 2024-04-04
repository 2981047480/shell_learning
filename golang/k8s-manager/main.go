package main

import (
	"example.com/m/v2/k8s-manager/handler/node"
	"example.com/m/v2/k8s-manager/handler/sys"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", sys.Ping)
	r.GET("/node/get", node.List)
	r.Run("127.0.0.1:8080") // 自己定义，不写默认是127.0.0.1:8080
}
