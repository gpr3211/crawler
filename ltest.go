package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

}
