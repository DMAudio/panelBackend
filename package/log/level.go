package log

func Debug(m ...interface{}) {
	logger.Debug(m...)
}

func Info(m ...interface{}) {
	logger.Info(m...)
}

func Warn(m ...interface{}) {
	logger.Warn(m...)
}

func Error(m ...interface{}) {
	logger.Error(m...)
}

func Panic(m ...interface{}) {
	logger.Panic(m...)
}

func Fatal(m ...interface{}) {
	logger.Fatal(m...)
}

func Print(m ...interface{}) {
	logger.Print(m...)
}

func Debugf(format string, m ...interface{}) {
	logger.Debugf(format, m...)
}

func Infof(format string, m ...interface{}) {
	logger.Infof(format, m...)
}

func Warnf(format string, m ...interface{}) {
	logger.Warnf(format, m...)
}

func Errorf(format string, m ...interface{}) {
	logger.Errorf(format, m...)
}

func Panicf(format string, m ...interface{}) {
	logger.Panicf(format, m...)
}

func Fatalf(format string, m ...interface{}) {
	logger.Fatalf(format, m...)
}

func Printf(format string, m ...interface{}) {
	logger.Infof(format, m...)
}

func With(v ...interface{}) Logger {
	return logger.With(v...)
}
