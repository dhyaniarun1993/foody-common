package logger

import (
	"context"
	"runtime"
	"strconv"

	"github.com/dhyaniarun1993/go-utility/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Configuration provides configuration for zap logger
type Configuration struct {
	Level  string `required:"true" split_words:"true"`
	Format string `required:"true" split_words:"true"`
}

// Logger provides structure logging backed by uber zap logger
type Logger struct {
	*zap.Logger
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	c := caller.FullPath()
	details := runtime.FuncForPC(caller.PC)
	if details != nil {
		c = details.Name() + ":" + strconv.Itoa(caller.Line)
	}
	enc.AppendString(c)
}

// CreateLogger creates logrus logger
func CreateLogger(configuration Configuration) *Logger {
	var logLevel zapcore.Level
	switch configuration.Level {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "WARN":
		logLevel = zapcore.WarnLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	default:
		panic("Invalid Log Level")
	}
	cfg := zap.Config{
		Encoding:         configuration.Format,
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: customCallerEncoder,

			StacktraceKey: "stacktrace",
		},
	}
	logger, _ := cfg.Build()
	return &Logger{logger}
}

// WithContext return a new Logger with as much context as possible
func (logger *Logger) WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		return logger
	}

	var newLogger *zap.Logger
	if ctxTraceID, ok := ctx.Value("traceId").(string); ok {
		newLogger = logger.With(zap.String("trace-id", ctxTraceID))
	}
	return &Logger{newLogger}
}

//WithError return a new Logger with Error context
func (logger *Logger) WithError(err error) *Logger {
	if err == nil {
		return logger
	}

	var newLogger *zap.Logger
	if appError, ok := errors.IsAppError(err); ok {
		newLogger = logger.With(zap.String("error", appError.Error()),
			zap.String("errorstack", appError.ErrorStack()))
	} else {
		newLogger = logger.With(zap.String("error", appError.Error()))
	}
	return &Logger{newLogger}
}
