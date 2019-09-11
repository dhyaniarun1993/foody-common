package logger

import (
	"context"
	"runtime"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dhyaniarun1993/foody-common/authentication"
	"github.com/dhyaniarun1993/foody-common/errors"
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
	if span := opentracing.SpanFromContext(ctx); span != nil {
		jaegerSpanContext := span.Context().(jaeger.SpanContext)
		traceID := jaegerSpanContext.TraceID().String()
		spanID := jaegerSpanContext.SpanID().String()
		newLogger = logger.With(zap.String("trace-id", traceID),
			zap.String("span-id", spanID))
	}

	if userID, ok := authentication.GetUserID(ctx); ok {
		if newLogger != nil {
			newLogger = newLogger.With(zap.String("user-id", userID))
		} else {
			newLogger = logger.With(zap.String("user-id", userID))
		}
	}

	if appID, ok := authentication.GetAppID(ctx); ok {
		if newLogger != nil {
			newLogger = newLogger.With(zap.String("app-id", appID))
		} else {
			newLogger = logger.With(zap.String("app-id", appID))
		}
	}

	if newLogger != nil {
		return &Logger{newLogger}
	}
	return logger
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
		newLogger = logger.With(zap.String("error", err.Error()))
	}
	return &Logger{newLogger}
}
