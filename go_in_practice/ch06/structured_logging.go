package ch06

import (
	"log/slog"
)

func main() {
	slog.Info("this is default logging")
	slog.Warn("Keep an eye on this, it might be an issue")
	slog.Error("oh no, an error happened here!")
	slog.Debug("this is good while developing")
}