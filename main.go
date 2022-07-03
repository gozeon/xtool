package main

import (
	_ "embed"
	"log"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

//go:embed icon/icon.png
var icon []byte

var links = map[string]string{
	"appName": "XTool",
	"github":  "https://github.com/gozeon",
	"twitter": "https://twitter.com/gozeonl",
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)

}

func main() {
	systray.Run(onReady, onExit)
}

func sysAlert(msg string) {
	beeep.Alert(links["appName"], msg, string(icon))
}

func onReady() {
	log.Println("onReady")
	systray.SetTemplateIcon(icon, icon)
	systray.SetTooltip(links["appName"])

	github := systray.AddMenuItem("Github", "gozeon")
	twitter := systray.AddMenuItem("Twitter", "gozeonl")

	alert := systray.AddMenuItem("Alert", "系统通知")
	clipboardMenu := systray.AddMenuItem("Clipboard", "粘贴板")

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-github.ClickedCh:
				open.Run(links["github"])
				log.Println("open", links["github"])
			case <-twitter.ClickedCh:
				open.Run(links["twitter"])
				log.Println("open", links["twitter"])
			case <-mQuit.ClickedCh:
				systray.Quit()
			case <-clipboardMenu.ClickedCh:
				cl, _ := clipboard.ReadAll()
				log.Println("clipboard", cl)
			case <-alert.ClickedCh:
				sysAlert("hello world")
			}
		}
	}()
}

func onExit() {
	log.Println("onExit")
}
