package storage_logger

import "fmt"

func (logger Logger) Debug(_ ...any)            {}
func (logger Logger) Debugf(_ string, _ ...any) {}

func (logger Logger) Info(_ ...any)            {}
func (logger Logger) Infof(_ string, _ ...any) {}

func (logger Logger) Warn(_ ...any)            {}
func (logger Logger) Warnf(_ string, _ ...any) {}

func (logger Logger) Error(args ...any) {
	_ = logger.write(logger.env, logger.service, "ERROR", build(args...))
}
func (logger Logger) Errorf(format string, args ...any) {
	_ = logger.write(logger.env, logger.service, "ERROR", fmt.Sprintf(format, args...))
}

func (logger Logger) Fatal(args ...any) {
	_ = logger.write(logger.env, logger.service, "FATAL", build(args...))
}
func (logger Logger) Fatalf(format string, args ...any) {
	_ = logger.write(logger.env, logger.service, "FATAL", fmt.Sprintf(format, args...))
}
