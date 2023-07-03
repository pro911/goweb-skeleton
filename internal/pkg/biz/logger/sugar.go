package logger

// Debug uses fmt.Sprint to construct and logger a message.
func Debug(args ...interface{}) {
	Sugar.Debug(args)
}

// Info uses fmt.Sprint to construct and logger a message.
func Info(args ...interface{}) {
	Sugar.Info(args)
}

// Warn uses fmt.Sprint to construct and logger a message.
func Warn(args ...interface{}) {
	Sugar.Warn(args)
}

// Error uses fmt.Sprint to construct and logger a message.
func Error(args ...interface{}) {
	Sugar.Error(args)
}

// DPanic uses fmt.Sprint to construct and logger a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	Sugar.DPanic(args)
}

// Panic uses fmt.Sprint to construct and logger a message, then panics.
func Panic(args ...interface{}) {
	Sugar.Panic(args)
}

// Fatal uses fmt.Sprint to construct and logger a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	Sugar.Fatal(args)
}

// Debugf uses fmt.Sprintf to logger a templated message.
func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args)
}

// Infof uses fmt.Sprintf to logger a templated message.
func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args, nil)
}

// Warnf uses fmt.Sprintf to logger a templated message.
func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args, nil)
}

// Errorf uses fmt.Sprintf to logger a templated message.
func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args, nil)
}

// DPanicf uses fmt.Sprintf to logger a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	Sugar.DPanicf(template, args, nil)
}

// Panicf uses fmt.Sprintf to logger a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	Sugar.Panicf(template, args, nil)
}

// Fatalf uses fmt.Sprintf to logger a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	Sugar.Fatalf(template, args, nil)
}
