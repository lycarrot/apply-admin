package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            //级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         //日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         //输出
	Director      string `mapstructure:"director" json:"director" yaml:"director"`                   //日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                //显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       //编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` //栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` //级别
	MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      //日志保留时间
}

func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel //调试目的的详细信息。
	case "info":
		return zapcore.InfoLevel //记录一般的操作信息或者应用程序状态
	case "warn":
		return zapcore.WarnLevel //记录一些可能需要注意的情况，但不是严重错误。
	case "error":
		return zapcore.ErrorLevel //记录错误，但程序仍然可以继续运行。
	case "dpanic":
		return zapcore.DPanicLevel //记录恐慌性错误，会导致程序崩溃之前的日志消息。
	case "pnaic":
		return zapcore.PanicLevel //恐慌的日志消息，并触发恐慌。
	case "fatal":
		return zapcore.FatalLevel //程序终止
	default:
		return zapcore.DebugLevel
	}
}
