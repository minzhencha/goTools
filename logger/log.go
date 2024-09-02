package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Level  zapcore.Level
	Syncer *lumberjack.Logger
}

const (
	Debug  = zapcore.DebugLevel
	Info   = zapcore.InfoLevel
	Warn   = zapcore.WarnLevel
	Error  = zapcore.ErrorLevel
	DPanic = zapcore.DPanicLevel
	Panic  = zapcore.PanicLevel
	Fatal  = zapcore.FatalLevel
)

func NewLogger(logs *Logger) (logger *zap.SugaredLogger) {
	// 日志输出设置
	logs = setLogger(logs)

	// 日志编码设置
	encoderCode := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(logs.Level)

	// 设置日志输出
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCode), zapcore.NewMultiWriteSyncer(zapcore.AddSync(logs.Syncer), zapcore.AddSync(os.Stdout)), atom)

	// 初始化日志
	logger = zap.New(core, zap.AddCaller(), zap.AddCaller()).Sugar()

	return logger
}

// 日志输出设置
func setLogger(logs *Logger) *Logger {
	// 日志文件设置
	if logs.Syncer.Filename == "" {
		logs.Syncer.Filename = "logs/" + time.Now().Format(time.DateOnly) + ".log"
	}
	// 日志文件大小设置
	if logs.Syncer.MaxSize == 0 {
		logs.Syncer.MaxSize = 1024
	}
	// 日志文件保留天数设置
	if logs.Syncer.MaxAge == 0 {
		logs.Syncer.MaxAge = 180
	}
	// 日志文件保留数量设置
	if logs.Syncer.MaxBackups == 0 {
		logs.Syncer.MaxBackups = 20
	}

	return logs
}
