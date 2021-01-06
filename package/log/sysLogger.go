package log

import (
	sysLog "log"
	"os"
)

type sysLogger struct{}

func (s *sysLogger) Fatal(v ...interface{}) {
	sysLog.Fatal(v...)
}

func (s *sysLogger) Fatalf(format string, v ...interface{}) {
	sysLog.Fatalf(format, v...)
}

func (s *sysLogger) Panic(v ...interface{}) {
	sysLog.Panic(v...)
}

func (s *sysLogger) Panicf(format string, v ...interface{}) {
	sysLog.Panicf(format, v...)
}

func (s *sysLogger) Print(v ...interface{}) {
	sysLog.Print(v...)
}

func (s *sysLogger) Printf(format string, v ...interface{}) {
	sysLog.Printf(format, v...)
}

func (s *sysLogger) Debug(v ...interface{}) {
	sysLog.Print(v...)
}

func (s *sysLogger) Debugf(format string, v ...interface{}) {
	sysLog.Printf(format, v...)
}

func (s *sysLogger) Info(v ...interface{}) {
	sysLog.Print(v...)
}

func (s *sysLogger) Infof(format string, v ...interface{}) {
	sysLog.Printf(format, v...)
}

func (s *sysLogger) Warn(v ...interface{}) {
	sysLog.Print(v...)
}

func (s *sysLogger) Warnf(format string, v ...interface{}) {
	sysLog.Printf(format, v...)
}

func (s *sysLogger) Error(v ...interface{}) {
	sysLog.Print(v...)
}

func (s *sysLogger) Errorf(format string, v ...interface{}) {
	sysLog.Printf(format, v...)
}

func (s *sysLogger) With(v ...interface{}) Logger {
	return &sysLogger{}
}

func (s *sysLogger) StdLogger() *sysLog.Logger {
	return sysLog.New(os.Stdout, "", 0)
}
