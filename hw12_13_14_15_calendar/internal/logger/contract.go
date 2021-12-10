package logger

type Contract interface {
	Debug(msg string)
	Warn(msg string)
	Info(msg string)
	Error(msg string)
}
