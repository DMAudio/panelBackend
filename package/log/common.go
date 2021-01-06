package log

var logger Logger

func init() {
	logger = &sysLogger{}
}

func SetLogger(l Logger) {
	logger = l
}
