package log

import (
	"fmt"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = ""
	logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Stack().
		CallerWithSkipFrameCount(3).
		Logger()
}

func Logger() zerolog.Logger {
	return logger
}

func Trace(msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Trace().Msg(msg)
}

func Debug(msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Debug().Msg(msg)
}

func Info(msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Info().Msg(msg)
}

func Warn(msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Warn().Msg(msg)
}

func Error(err error, msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Error().Stack().Err(err).Msg(msg)
	fmt.Printf("%s", debug.Stack())
}

func Fatal(err error, msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Fatal().Stack().Err(err).Msg(msg)
	fmt.Printf("%s", debug.Stack())
}

func Panic(err error, msgTemplate string, v ...any) {
	msg := fmt.Sprintf(msgTemplate, v...)
	logger.Panic().Stack().Err(err).Msg(msg)
	fmt.Printf("%s", debug.Stack())
}
