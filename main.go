package main

import (
	_ "embed"
	"log"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

//go:embed icon/icon.png
var icon []byte

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	systray.Run(onReady, onExit)
}

func onReady() {
	log.Println("onReady")
	systray.SetTemplateIcon(icon, icon)
	// systray.SetTitle("xtool")
	systray.SetTooltip("xtool")

	github := systray.AddMenuItem("Github", "gozeon")
	twitter := systray.AddMenuItem("Twitter", "gozeonl")

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-github.ClickedCh:
				log.Println("open", "https://github.com/gozeon")
				open.Run("https://github.com/gozeon")
			case <-twitter.ClickedCh:
				log.Println("open", "https://twitter.com/gozeonl")
				open.Run("https://twitter.com/gozeonl")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	log.Println("onExit")
}
