package logger

import (
	"github.com/pkg/errors"
	zapLog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var mapLogLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func NewLogger(config *Config) (*Logger, error) {
	logger := &Logger{
		Config: config,
	}
	err := logger.InitLogger()
	return logger, err
}

func (l *Logger) InitLogger() error {
	if l.Config == nil {
		return errors.New("param error : config is nil")
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(l.Config.GetZapEncoderConfig()),                                    // 编码
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(l.Config.GetHook())), // 打印输出端
		GetZapLevel(l.Config.LogLevel, zapcore.DebugLevel),                                           // 打印日志级别                            // 日志
	)

	// 开启开发模式，调用跟踪
	caller := zapLog.AddCaller()
	// 开启堆栈跟踪
	stacktrace := zapLog.AddStacktrace(GetZapLevel(l.Config.StacktraceLevel, zapcore.ErrorLevel))
	// 开启文件及行号
	development := zapLog.Development()
	// 设置初始化字段
	filed := zapLog.Fields(zapLog.String("serviceName", l.Config.ServiceName))

	// 构建日志
	l.log = zapLog.New(core, caller, stacktrace, filed, development)
	l.logSugar = l.log.Sugar()
	return nil
}

func GetZapLevel(levelKey string, defaultLevel zapcore.Level) zapcore.Level {
	levelKey = strings.ToLower(levelKey)
	result, ok := mapLogLevel[levelKey]
	if ok {
		return result
	}
	return defaultLevel
}
