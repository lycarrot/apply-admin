package internal

import (
	"gin-pro/global"
	"go.uber.org/zap/zapcore"
	"os"
)

type fileRotatelogs struct {
}

var FileRotatelogs = new(fileRotatelogs)

// zapcore.NewMultiWriteSyncer 是一个函数，用于创建多个写入同步器,它接受一系列的写入同步器作为参数，并返回一个新的写入同步器，该同步器会将日志同时写入到所有提供的写入目标中。
// zapcore.AddSync 是用于创建写入同步器的函数，它接受一个实现了 io.Writer 接口的对象作为参数，并返回一个对应的写入同步器。
// os.Stdout 标准输出流，是一个实现了 io.Writer 接口的对象。
func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.GVA_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.GVA_CONFIG.Zap.LogInConsole {
		//将日志写入到文件和控制台输出
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
