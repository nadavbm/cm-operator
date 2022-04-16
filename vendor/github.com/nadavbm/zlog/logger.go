package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func InitLogger() *Logger {
	encfg := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.CapitalLevelEncoder,
	}

	options := zapcore.NewCore(zapcore.NewJSONEncoder(encfg), os.Stdout, zap.DebugLevel)

	l := zap.New(options)

	logger := &Logger{
		Logger: l,
	}
	return logger
}
