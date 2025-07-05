package tests

import "log/slog"

func NewTestLogger() *slog.Logger {
	return slog.New(slog.DiscardHandler)
}
