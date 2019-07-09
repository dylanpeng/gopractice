package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
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
