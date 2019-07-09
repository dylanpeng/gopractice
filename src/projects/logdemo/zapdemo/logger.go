package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func main() {
	LumberjackLog()
}

func QuickLoggerDemo() {
	url := "http://www.google.com"

	//print format log
	loggerDev, _ := zap.NewDevelopment()
	defer loggerDev.Sync()
	loggerDev.Info("loggerDev failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	//sugar 可以不指定类型，代码友好但效率比较低
	sugarDev := loggerDev.Sugar()
	sugarDev.Warnw("sugarDev failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarDev.Warnf("sugarDev Failed to fetch URL: %s \n", url)

	//print json log
	loggerPro, _ := zap.NewProduction()
	defer loggerPro.Sync()
	loggerPro.Info("loggerPro failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugarPro := loggerPro.Sugar()
	sugarPro.Errorw("sugarPro failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarPro.Errorf("sugarPro Failed to fetch URL: %s", url)
}

func ConfigLogDemo() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "console",                                           // 输出格式 console json
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器
		OutputPaths:      []string{"stdout", "./logs/sample.log"},             // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	//构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	logger.Info("log 初始化成功")

	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Error("log error")

}

func LumberjackLog() {
	hook := lumberjack.Logger{
		Filename:   "./logs/lumberjacklog.log", // 日志文件路径
		MaxSize:    1,                          // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 300,                        // 日志文件最多保存多少个备份
		MaxAge:     7,                          // 文件最多保存多少天
		Compress:   false,                      // 是否压缩
		LocalTime:  true,                       // 是否用当地时间命名文件
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                        // 编码
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印
		atomicLevel,                                                                     // 日志
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构建日志
	logger := zap.New(core, caller, development, filed)

	var count int
	for {
		logger.Info("log 初始化成功")
		logger.Info("无法获取网址",
			zap.String("url", "http://www.baidu.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second))

		count++
		if count > 10000 {
			break
		}
	}
}
