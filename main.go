package main

import (
	"log"
	"os"

	"github.com/zserge/webview"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Incorrect number of parameters, usage: program url\n")
		os.Exit(1)
	}

	w := webview.New(webview.Settings{
		Debug:                  true,
		ExternalInvokeCallback: nil,
		Width:                  800,
		Height:                 600,
		Resizable:              true,
		Title:                  "Dashboard",
		URL:                    os.Args[1],
	})

	w.SetFullscreen(true)
	log.Printf("Starting for url: %v\n", os.Args[1])
	w.Run()
}
