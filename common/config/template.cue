import ( "time")
time_duration: int64

_serviceBase: {
    shutdownTimeout: time_duration | *time.ParseDuration("24h")
}