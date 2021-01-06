package libzap

import (
	sysLog "log"

	"go.uber.org/zap"

	"github.com/DMAudio/panelBackend/package/log"
)

type zapLogger struct {
	*zap.SugaredLogger
}

func (l zapLogger) Print(v ...interface{}) {
	l.Info(v...)
}
func (l zapLogger) Printf(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func (l zapLogger) With(v ...interface{}) log.Logger {
	return &zapLogger{l.SugaredLogger.With(v...)}
}

func (l zapLogger) StdLogger() *sysLog.Logger {
	return zap.NewStdLog(l.SugaredLogger.Desugar())
}
