package main

import (
	"base-showcase/core"
	"base-showcase/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//import "github.com/gin-gonic/gin"

/**
参考来源
https://mp.weixin.qq.com/s/t1ngDeT5Jq1ueixDWT9Q2w
https://juejin.cn/post/7213297003869569081

完整的代码：https://github.com/ian-kevin126/gin_common_web_server
*/

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	fmt.Println("hello world")
	gin.SetMode(AppMode)

	// TODO：1.配置初始化
	global.EWA_VIPER = core.InitializeViper()
	// TODO：2.日志
	global.EWA_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.EWA_LOG)
	global.EWA_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))
	// TODO：3.数据库连接
	// TODO：4.其他初始化
	// TODO：5.启动服务

	r := gin.Default()

	//测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":8080")

}
