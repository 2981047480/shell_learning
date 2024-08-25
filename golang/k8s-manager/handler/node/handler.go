package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义list，获取NodeList的接口
func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "node list",
	})
}
// 有了List，我们还需要定义一个路由来暴露这个方法 

// // 定义获取NodeList的接口
// type GetNodeListHandler interface {
// 	GetNodeList() (ctx *gin.Context)
// }

// type kubeClient interface {
// 	// 定义k8s客户端的接口
// 	GetNodeListHandler
// }

// func main() {
// 	gin
// }
