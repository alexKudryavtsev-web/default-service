package logger

import (
	"fmt"
	"os"
	"strings"

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

func New(level string) (*Logger, error) {
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

 skipFrameCount := 3
 logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

 return &Logger{
  logger: &logger,
 }, nil
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
 l.msg("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
 l.msg("info", message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
 l.msg("warn", message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
 l.msg("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
 l.msg("fatal", message, args...)
 os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
 l.logger.Info().Msgf(message, args...)
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
 switch msg := message.(type) {
 case error:
  l.log(msg.Error(), args...)
 case string:
  l.log(msg, args...)
 default:
  l.log(fmt.Sprintf("%s message %v (type: %T) has unknown type", level, message, message), args...)
 }
}