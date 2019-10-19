package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/zserge/webview"
)

func main() {
	if len(os.Args) != 4 {
		log.Printf("Incorrect number of parameters, usage: program url width height\n")
		os.Exit(1)
	}

	width, _ := strconv.ParseInt(os.Args[2], 10, 64)
	height, _ := strconv.ParseInt(os.Args[3], 10, 64)

	w := webview.New(webview.Settings{
		Debug:                  true,
		ExternalInvokeCallback: nil,
		Width:                  int(width),
		Height:                 int(height),
		Resizable:              true,
		Title:                  "Dashboard",
		URL:                    os.Args[1],
	})

	if strings.ToLower(os.Getenv("DASHBOARD_DISABLE_FULLSCREEN")) != "y" && strings.ToLower(os.Getenv("DASHBOARD_DISABLE_FULLSCREEN")) != "yes" {
		w.SetFullscreen(true)
	}

	log.Printf("Starting for url: %v\n", os.Args[1])
	w.Run()
}
