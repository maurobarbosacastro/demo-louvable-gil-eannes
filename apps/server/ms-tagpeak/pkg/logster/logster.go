package logster

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CustomLogster struct {
	zerolog.Logger
}

var MiddlewareLogster CustomLogster
var wrapperLogster CustomLogster

// This package is a wrapper for zerolog

func InitLogster(currentEnv string, loggerLevel string) {
	zerolog.SetGlobalLevel(getLogLevel(loggerLevel))

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	//Logger for the rest of the app has more option
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	// Middleware logger is as simple as possible
	MiddlewareLogster = CustomLogster{log.Logger}
	MiddlewareLogster.Logger = zerolog.New(output).With().Timestamp().Logger()

	//Skip 3 frames (logster -> event -> logster) to get to the original caller.
	//The reason for needing 3 frames is because of how the call stack looks:
	//1-The actual zerolog logging code
	//2-The zerolog event creation
	//3-Your wrapper function
	//4-The actual caller's code (this is what you want to show)
	log.Logger = zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
	wrapperLogster.Logger = zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(3).Logger()

	if currentEnv == "dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Info("Dev Logger level: 0")
	}
}

func Print(msg string) {
	log.Trace().Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

// Fatal will stop the program
func Fatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}

// Panic will stop the program and print a stack. Last resort method.
func Panic(err error, msg string) {
	log.Panic().Err(err).Msg(msg)
}

func StartFuncLog() {
	wrapperLogster.Info().Msg(fmt.Sprintf("Start %s", getCallerFunctionName()))
}
func StartFuncLogMsg(msg interface{}) {
	wrapperLogster.Info().Msg(fmt.Sprintf("Start %s - %s", getCallerFunctionName(), msg))
}

func EndFuncLog() {
	wrapperLogster.Info().Msg(fmt.Sprintf("End %s", getCallerFunctionName()))
}

func EndFuncLogMsg(msg interface{}) {
	wrapperLogster.Info().Msg(fmt.Sprintf("End %s - %s", getCallerFunctionName(), msg))
}

func getCallerFunctionName() string {
	// Skip 1 tells runtime.Caller to get the frame of the calling function
	// rather than this function itself
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	// Get function details from program counter
	details := runtime.FuncForPC(pc)
	if details == nil {
		return "unknown"
	}

	return details.Name()
}

func LogsterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log the request
		MiddlewareLogster.Info().Fields(map[string]interface{}{
			"Method":    c.Request().Method,
			"URL":       c.Request().URL.Path,
			"Query":     c.Request().URL.RawQuery,
			"Host":      c.Request().Host,
			"IP":        c.RealIP(),
			"Status":    c.Response().Status,
			"UserAgent": c.Request().UserAgent(),
		}).Msg("Request")

		// call the next middleware/handler
		err := next(c)
		if err != nil {
			MiddlewareLogster.Error().Fields(map[string]interface{}{
				"error": err.Error(),
			}).Msg("Response")
			return err
		}

		return nil
	}
}

func getLogLevel(logLevel string) zerolog.Level {
	nivel, _ := strconv.Atoi(logLevel)
	var level zerolog.Level

	switch nivel {
	case -1:
		level = zerolog.TraceLevel
	case 0:
		level = zerolog.DebugLevel
	case 1:
		level = zerolog.InfoLevel
	case 2:
		level = zerolog.WarnLevel
	case 3:
		level = zerolog.ErrorLevel
	case 4:
		level = zerolog.FatalLevel
	case 5:
		level = zerolog.PanicLevel
	default:
		level = zerolog.InfoLevel
	}

	return level
}
