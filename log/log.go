package log

import "fmt"

func Infof(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	log.logger.Info(msg)
}

func Errorf(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	log.logger.Error(msg)
}

func Debugf(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	log.logger.Debug(msg)
}

func Warnf(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	log.logger.Warn(msg)
}

func Fatalf(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	log.logger.Fatal(msg)
}
