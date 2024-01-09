package logger

func (logger *Logger) Debug(args ...any) {
	logger.printLog(func() {
		logger.sugar.Debug(args...)
	})
}

func (logger *Logger) Debugf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Debugf(template, args...)
	})
}

func (logger *Logger) Info(args ...any) {
	logger.printLog(func() {
		logger.sugar.Info(args...)
	})
}

func (logger *Logger) Infof(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Infof(template, args...)
	})
}

func (logger *Logger) Warn(args ...any) {
	logger.printLog(func() {
		logger.sugar.Warn(args...)
	})
}

func (logger *Logger) Warnf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Warnf(template, args...)
	})
}

func (logger *Logger) Error(args ...any) {
	logger.printLog(func() {
		logger.sugar.Error(args...)
	})
}

func (logger *Logger) Errorf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Errorf(template, args...)
	})
}

func (logger *Logger) Fatal(args ...any) {
	logger.printLog(func() {
		logger.sugar.Fatal(args...)
	})
}

func (logger *Logger) Fatalf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Fatalf(template, args...)
	})
}
