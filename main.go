package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"image/png"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/slogcolor"
	"github.com/ncruces/zenity"
)

func setIcon(wnd *giu.MasterWindow, fsys fs.FS) error {
	f, err := fsys.Open("icon.png")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	img, err := png.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	wnd.SetIcon(img)

	return nil
}

// getLogLevel gets the log level from command-line flags.
func getLogLevel() slog.Leveler {
	switch s := strings.ToLower(*logLevelFlag); s {
	case "":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic(fmt.Sprintf("invalid log level: \"%s\"; should be one of \"debug\", \"info\", \"warn\", \"error\"", s))
	}
}

func loop() {
	giu.SingleWindow().Layout(giu.Label("horalky"))
}

func main() {
	// Logger
	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))

	slog.Info("GCharted version " + Version)
	slog.Info("Go version " + runtime.Version())

	slog.Debug("creating window")
	wnd := giu.NewMasterWindow("GCharted", 1280, 720, 0)
	slog.Debug("setting window icon")
	if err := setIcon(wnd, assets.FS); err != nil {
		slog.Error("Failed to set window icon", "err", err)
		zenity.Error("Failed to set window icon")
	}
	slog.Info("showing window")
	wnd.Run(loop)
}
