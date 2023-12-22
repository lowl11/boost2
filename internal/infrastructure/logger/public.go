package logger

func (logger Logger) Debug(args ...any) {
	logger.sugar.Debug(args...)
	logger.storageLogger.Debug(args...)
}

func (logger Logger) Debugf(template string, args ...any) {
	logger.sugar.Debugf(template, args...)
	logger.storageLogger.Debugf(template, args...)
}

func (logger Logger) Info(args ...any) {
	logger.sugar.Info(args...)
	logger.storageLogger.Info(args...)
}

func (logger Logger) Infof(template string, args ...any) {
	logger.sugar.Infof(template, args...)
	logger.storageLogger.Infof(template, args...)
}

func (logger Logger) Warn(args ...any) {
	logger.sugar.Warn(args...)
	logger.storageLogger.Warn(args...)
}

func (logger Logger) Warnf(template string, args ...any) {
	logger.sugar.Warnf(template, args...)
	logger.storageLogger.Warnf(template, args...)
}

func (logger Logger) Error(args ...any) {
	logger.sugar.Error(args...)
	logger.storageLogger.Error(args...)
}

func (logger Logger) Errorf(template string, args ...any) {
	logger.sugar.Errorf(template, args...)
	logger.storageLogger.Errorf(template, args...)
}

func (logger Logger) Fatal(args ...any) {
	logger.sugar.Fatal(args...)
	logger.storageLogger.Fatal(args...)
}

func (logger Logger) Fatalf(template string, args ...any) {
	logger.sugar.Fatalf(template, args...)
	logger.storageLogger.Fatalf(template, args...)
}
