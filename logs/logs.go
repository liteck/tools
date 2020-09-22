package logs

import (
	"go.uber.org/zap"
)

var Suger  *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer func() {
		_=logger.Sync() // flushes buffer, if any
	}()
	Suger = logger.Sugar()
}

func Error(args ...interface{}) {
	Suger.Error(args)
}