package core

import (
	"base-showcase/core/internal"
	"base-showcase/global"
	"base-showcase/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.EWA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.EWA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.EWA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.EWA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
