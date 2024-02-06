package utils

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var baseOnce sync.Once
var once sync.Once
var baseLog zerolog.Logger
var Log zerolog.Logger

func GetBaseLog() *zerolog.Logger {
	baseOnce.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.TraceLevel)
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FieldsExclude: []string{
				"user_agent",
				"git_revision",
				"go_version",
			},
		}

		if os.Getenv("APP_ENV") != "development" {
			filelogger := &lumberjack.Logger{
				Filename:   "share_room.log",
				MaxSize:    5,
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, filelogger)
		}

		baseLog = zerolog.New(output).Level(zerolog.Level(logLevel)).With().Timestamp().Logger()
	})
	return &baseLog
}

func GetLog() *zerolog.Logger {
	once.Do(func() {
		var gitRevision string
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		Log = GetBaseLog().With().Caller().
			Str("git_revision", gitRevision).Str("go_version", buildInfo.GoVersion).Logger()
		zerolog.DefaultContextLogger = &Log
	})
	return &Log
}
