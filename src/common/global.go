package common

import "gopractice/lib/logger"

var Logger *logger.Logger

func InitLogger(config *logger.Config) (err error) {
	Logger, err = logger.NewLogger(config)
	return err
}
