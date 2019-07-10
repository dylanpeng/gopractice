package logger

import (
	zapLog "go.uber.org/zap"
)

type Logger struct {
	Config   *Config
	log      *zapLog.Logger
	logSugar *zapLog.SugaredLogger
}

func (l *Logger) Debug(msg string, fields ...zapLog.Field) {
	l.log.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zapLog.Field) {
	l.log.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zapLog.Field) {
	l.log.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapLog.Field) {
	l.log.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zapLog.Field) {
	l.log.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zapLog.Field) {
	l.log.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zapLog.Field) {
	l.log.Fatal(msg, fields...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logSugar.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.logSugar.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logSugar.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logSugar.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logSugar.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logSugar.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logSugar.Fatalf(template, args...)
}
