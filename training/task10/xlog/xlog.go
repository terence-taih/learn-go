package xlog

import (
	"context"
	"gitlab.myteksi.net/gophers/go/commons/util/log/logging"
	"gitlab.myteksi.net/gophers/go/commons/util/log/structlog"
)

func XStructLog() {
	logging.New(&logging.Config{
		Level:                 3,
		Tag:                   "Thai-Demo",
		DisableStackWhenError: true,
	})
	logObj := structlog.New("LogObj").With("corrlationId", "1234")
	ctx := structlog.NewContextWithLogger(context.Background(), logObj)

	structlog.FromContext(ctx, "LogObj").Error("Thai-Demo", "Testing error %s", "OK")
	logging.ShutdownAndClose()
}

func XCommonLog() {
	logTag := "Thai-Demo"
	logCfg := logging.Config{
		Level:                 3,
		Tag:                   "Thai-Demo",
		DisableStackWhenError: true,
	}

	logger := logging.New(&logCfg)
	logger.Debug(logTag, "Testing debug: %s", "debug params")
	logger.Error(logTag, "Testing error: %s", "error params")
	logging.ShutdownAndClose()

}
