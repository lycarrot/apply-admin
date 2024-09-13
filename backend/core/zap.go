package core

import (
	"gin-pro/core/internal"
	"gin-pro/global"
	"gin-pro/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// zapcore.NewTee 用于创建一个多路复用的日志核心。它接受多个 zapcore.Core 实例作为参数，并返回一个新的核心 ，
// 可以同时将日志消息发送到这些核心。这对于将日志同时记录到多个目的地非常有用，比如同时记录到文件和网络中。
//zap.New 是创建 Zap 日志记录器（logger）的函数。它接受一个或多个 zapcore.Core 实例作为参数，并返回一个新的 Logger。

// zapcore.NewTee 用于创建一个将日志消息发送到多个核心的核心。
// zap.New 用于创建一个新的 Logger 实例。
// zapcore.NewCore 用于创建一个单个日志核心，决定了日志消息的编码、写入目的地和记录级别。
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok {
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.GVA_CONFIG.Zap.ShowLine {
		//WithOptions 方法用于创建一个基于现有日志记录器的新日志记录器，而 AddCaller 选项则用于启用记录日志调用方（caller）信息的功能。
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
