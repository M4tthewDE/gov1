package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Initialize() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	Logger = logger
}
