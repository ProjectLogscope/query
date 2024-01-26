package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/hardeepnarang10/query/cmd/api/internal/config"
	"github.com/hardeepnarang10/query/pkg/logger"
	"github.com/hardeepnarang10/query/pkg/writer"
)

var mapLogLevel map[string]slog.Level = map[string]slog.Level{
	config.LogLevelDebug: slog.LevelDebug,
	config.LogLevelWarn:  slog.LevelInfo,
	config.LogLevelInfo:  slog.LevelWarn,
	config.LogLevelError: slog.LevelError,
}

func initLogger(logPath string, logLevel string, addSourceInfo bool) {
	lvl := mapLogLevel[logLevel]
	outputWriter := io.MultiWriter(os.Stdout, writer.New(logPath))
	l := logger.New(outputWriter, logger.TypeJSON,
		&slog.HandlerOptions{
			AddSource: addSourceInfo,
			Level:     lvl,
		},
	)
	slog.SetDefault(slog.New(l))
}
