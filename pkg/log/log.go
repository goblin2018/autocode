package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.SugaredLogger
}

var L *Logger

func Init(mode string, level string) {
	writeSyner := getWriter(mode)
	encoder := getEncoder(mode)
	core := zapcore.NewCore(encoder, writeSyner, getLogLevel(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	L = &Logger{
		logger.Sugar(),
	}
}

func getLogLevel(level string) zapcore.LevelEnabler {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

func getEncoder(mode string) zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	if mode == "dev" {
		ec = zap.NewProductionEncoderConfig()
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		ec.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	ec.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.000")
	return zapcore.NewConsoleEncoder(ec)
}

func getWriter(mode string) zapcore.WriteSyncer {

	lw := &lumberjack.Logger{
		Filename:   "logs/main.log",
		MaxSize:    100,
		MaxAge:     60,
		MaxBackups: 2,
	}

	if mode == "dev" {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lw), zapcore.AddSync(os.Stdout))
	} else {
		return zapcore.AddSync(lw)
	}
}

func Debug(args ...interface{}) {
	L.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	L.Debugf(template, args...)
}

func Info(args ...interface{}) {
	L.Info(args...)
}

func Infof(template string, args ...interface{}) {
	L.Infof(template, args...)
}

func Warn(args ...interface{}) {
	L.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	L.Warnf(template, args...)
}

func Error(args ...interface{}) {
	L.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	L.Errorf(template, args...)
}
