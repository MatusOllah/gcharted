package main

import (
	"log/slog"
	"os"
	"runtime"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/slogcolor"
)

func loop() {
	giu.SingleWindow().Layout(giu.Label("horalky"))
}

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, nil)))

	slog.Info("GCharted version " + Version)
	slog.Info("Go version " + runtime.Version())

	wnd := giu.NewMasterWindow("GCharted", 1280, 720, 0)
	wnd.Run(loop)
}
