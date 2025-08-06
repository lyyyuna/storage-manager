package log

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

func newLogger(skip int) *Logger {
	var w io.Writer
	w = os.Stdout
	consoleDebugging := zapcore.AddSync(w)

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core
	core = zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(zapcore.DebugLevel))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(skip))

	return &Logger{logger: logger}
}

var g *Logger = newLogger(2)
var s *Logger = newLogger(2)

// NewWithTag 生成一个新 logger，并带上自定义标签
func NewWithTag(key, val string) *Logger {
	return &Logger{
		logger: s.logger.With(zap.String(key, val)),
	}
}

func (l *Logger) WithString(key, val string) *Logger {
	logger := l.logger.With(zap.String(key, val))

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Sugar().Errorf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Sugar().Error(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Sugar().Fatalf(format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Sugar().Fatal(v...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Sugar().Panicf(format, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.logger.Sugar().Panic(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Sugar().Info(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Sugar().Info(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger.Sugar().Debugf(format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger.Sugar().Debug(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger.Sugar().Warnf(format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.logger.Sugar().Warn(v...)
}

func (l *Logger) Print(v ...interface{}) {
	l.logger.Sugar().Info(v...)
}

func Fatalf(format string, v ...interface{}) {
	g.Fatalf(format, v...)
}

func Fatal(v ...interface{}) {
	g.Fatal(v...)
}

func Errorf(format string, v ...interface{}) {
	g.Errorf(format, v...)
}

func Error(v ...interface{}) {
	g.Error(v...)
}

func Warnf(format string, v ...interface{}) {
	g.Warnf(format, v...)
}

func Warn(v ...interface{}) {
	g.Warn(v...)
}

func Panicf(format string, v ...interface{}) {
	g.Panicf(format, v...)
}

func Panic(v ...interface{}) {
	g.Panic(v...)
}

func Infof(format string, v ...interface{}) {
	g.Infof(format, v...)
}

func Info(v ...interface{}) {
	g.Info(v...)
}

func Debugf(format string, v ...interface{}) {
	g.Debugf(format, v...)
}

func Debug(v ...interface{}) {
	g.Debug(v...)
}

func Print(v ...interface{}) {
	g.Print(v...)
}

func Printf(format string, v ...interface{}) {
	g.Printf(format, v...)
}
