package libzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/DMAudio/panelBackend/package/errgo/errors"
	"github.com/DMAudio/panelBackend/package/log"
)

var ErrCantInitLogger = errors.New("can't initialize logger")

type traceLevel struct{}

func (traceLevel) Enabled(lvl zapcore.Level) bool {
	switch lvl {
	case zapcore.DPanicLevel,
		zapcore.PanicLevel,
		zapcore.FatalLevel:
		return true
	}
	return false
}

func Setup() (d func() error) {
	var err error
	var logger *zap.Logger
	// switch {
	// case runtime.RunAsDevelop():
	logger, err = zap.NewDevelopment(zap.AddCallerSkip(1), zap.AddStacktrace(traceLevel{}))
	// case runtime.RunAsTesting():
	// 	logger = zap.NewExample(zap.AddCallerSkip(1))
	// case runtime.RunAsProduct():
	// 	logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	// default:
	// 	err = errors.New("invalid runtime level")
	// }
	if err != nil {
		err = errors.Because(ErrCantInitLogger, err, "")
		log.Fatal(err)
	}
	log.SetLogger(zapLogger{logger.Sugar()})
	return logger.Sync
}
