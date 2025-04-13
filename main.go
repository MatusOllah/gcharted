package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"runtime"

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

func loop() {
	giu.SingleWindow().Layout(giu.Label("horalky"))
}

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, nil)))

	slog.Info("GCharted version " + Version)
	slog.Info("Go version " + runtime.Version())

	wnd := giu.NewMasterWindow("GCharted", 1280, 720, 0)
	if err := setIcon(wnd, assets.FS); err != nil {
		slog.Error("Failed to set window icon", "err", err)
		zenity.Error("Failed to set window icon")
	}
	wnd.Run(loop)
}
