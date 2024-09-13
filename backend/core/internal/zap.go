package internal

import (
	"gin-pro/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type _zap struct {
}

var Zap = new(_zap)

func (z *_zap) GetEncoder() zapcore.Encoder {
	//日志消息以 JSON 对象的形式进行编码,日志聚合系统或日志存储器，比如 Elasticsearch、Logstash、Kibana（ELK 栈）等。
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	//日志消息格式化为易于阅读的文本形式，并输出到控制台,适合控制台输出
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// 用于配置日志编码器的行为，定制化日志输出的格式和样式。
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GVA_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	// 将日志输出到文件或控制台
	writer := FileRotatelogs.GetWriteSyncer(l.String())
	//创建一个定制的日志核心，并将其传递给 Zap 的 Logger，从而实现灵活的日志记录配置。
	//encoder：日志编码器，用于将日志消息编码成特定的格式，例如 JSON、控制台格式等。
	//ws：写入同步器（write syncer），用于指定日志消息的输出目的地，可以是文件、控制台、网络等。
	//level：日志级别启用器（level enabler），用于确定哪些级别的日志消息将被记录下来。
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GVA_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// 根据配置文件的日志等级决定日志数据是否收集
// 错误日志等级:
// DebugLevel：对应值为 -1
// InfoLevel：对应值为 0
// WarnLevel：对应值为 1
// ErrorLevel：对应值为 2
// DPanicLevel：对应值为 3
// PanicLevel：对应值为 4
// FatalLevel：对应值为 5
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GVA_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// 根据配置文件level获取对应的zap.LevelEnablerFunc,用于确定是否启用特定日志级别的函数签名。
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { return level == zap.DebugLevel }
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { return level == zap.InfoLevel }
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { return level == zap.WarnLevel }
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { return level == zap.ErrorLevel }
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { return level == zap.DPanicLevel }
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { return level == zap.PanicLevel }
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { return level == zap.FatalLevel }
	default:
		return func(level zapcore.Level) bool { return level == zap.DebugLevel }

	}
}
