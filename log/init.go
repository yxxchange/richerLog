package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var defaultEncoderConfig = zapcore.EncoderConfig{
	MessageKey:     "msg",
	LevelKey:       "level",
	TimeKey:        "time",
	CallerKey:      "caller",
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     localTimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     zapcore.FullNameEncoder,
}

func init() {
	log = &Log{
		config:      defaultEncoderConfig,
		encoderMode: "json",
		logMode:     zapcore.InfoLevel,
	}
}

var log *Log

type Log struct {
	config zapcore.EncoderConfig
	logger *zap.Logger

	encoderMode string // json or console
	logMode     zapcore.Level
	logWriter   []io.Writer
}

func WithJsonEncoder(cfg zapcore.EncoderConfig) {
	log.encoderMode = "json"
	log.config = cfg
}

func WithLogWriter(writers ...io.Writer) {
	log.logWriter = writers
}

func WithConsoleEncoder(cfg zapcore.EncoderConfig) {
	log.encoderMode = "console"
	log.config = cfg
}

func Init() {
	var (
		encoder zapcore.Encoder
		cores   []zapcore.Core
	)
	encoder = initEncoder()
	cores = initCore(encoder)

	log.logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}

func initEncoder() zapcore.Encoder {
	switch log.encoderMode {
	case "json":
		return zapcore.NewJSONEncoder(log.config)
	case "console":
		return zapcore.NewConsoleEncoder(log.config)
	default:
		panic(fmt.Sprintf("unknown encoder mode: %s", log.encoderMode))
	}
}

func initCore(encoder zapcore.Encoder) []zapcore.Core {
	var cores []zapcore.Core
	if len(log.logWriter) == 0 {
		log.logWriter = append(log.logWriter, os.Stdout)
	}
	for _, writer := range log.logWriter {
		core := zapcore.NewCore(encoder, zapcore.AddSync(writer), log.logMode)
		cores = append(cores, core)
	}
	return cores
}

func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
