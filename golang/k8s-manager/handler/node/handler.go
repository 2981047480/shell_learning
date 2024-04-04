package node

import (
	"context"
	"net/http"

	"example.com/m/v2/k8s-manager/handler/common"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 定义list，获取NodeList的接口
func List(ctx *gin.Context) {
	// 首先配置kubeconfig，得到client对象
	clientSet, err := common.GetClientSet()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    10001,
			"message": err.Error(),
		})
		return
	}
	// 发出获取节点列表的请求
	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    10003,
			"message": err.Error(),
		})
		return
	}
	// 得到k8s cluster的apiserver的响应，处理这个响应的返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": nodeList,
	})
	// 给前端做响应

}

// ctx.JSON(http.StatusOK, gin.H{
// 	"code": 200,
// 	"data": "node list",
// })

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
