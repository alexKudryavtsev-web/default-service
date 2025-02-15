package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

func New(level string, destination string) (*Logger, error) {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		return nil, fmt.Errorf("invalid log level: %s", level)
	}

	zerolog.SetGlobalLevel(l)

	var output io.Writer

	switch strings.ToLower(destination) {
	case "file":
		logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("error opening log file: %w", err)
		}
		output = logFile
	case "console":
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		}
		output = consoleWriter
	default:
		return nil, fmt.Errorf("invalid log destination: %s. Use 'file' or 'console'", destination)
	}

	skipFrameCount := 3
	logger := zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}, nil
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.log("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log("info", message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log("warn", message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.log("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.log("fatal", message, args...)
	os.Exit(1)
}

func (l *Logger) log(level string, message interface{}, args ...interface{}) {
	var event *zerolog.Event
	switch level {
	case "debug":
		event = l.logger.Debug()
	case "info":
		event = l.logger.Info()
	case "warn":
		event = l.logger.Warn()
	case "error":
		event = l.logger.Error()
	case "fatal":
		event = l.logger.Fatal()
	default:
		event = l.logger.Info()
	}

	switch msg := message.(type) {
	case error:
		event.Msgf(msg.Error(), args...)
	case string:
		event.Msgf(msg, args...)
	default:
		event.Msgf(fmt.Sprintf("level message %v (type: %T) has unknown type", message, message), args...)
	}
}
