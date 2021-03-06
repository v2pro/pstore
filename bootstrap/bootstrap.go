package bootstrap

import (
	"github.com/v2pro/lego/zap_adatper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/v2pro/plz/log"
	"github.com/v2pro/plz/server"
	"github.com/v2pro/lego/echo_adapter"
	_ "github.com/v2pro/lego/jsoniter_adapter"
)

func init() {
	logger, _ := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console",
		DisableStacktrace: true,
		DisableCaller: true,
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   nil,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	log.AddLoggerProvider(func(loggerKv []interface{}) log.Logger {
		return zap_adatper.Adapt(logger)
	})
	server.AddServerExecutor(echo_adapter.StartServer)
}
