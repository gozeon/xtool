package main

import (
	_ "embed"
	"log"
	"xtool/icon"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/nleeper/goment"
	"github.com/skratchdot/open-golang/open"
)

//go:embed icon/icon.png
var iconM []byte

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
	beeep.Alert(links["appName"], msg, string(iconM))
}

func getTime(format string) string {
	g, err := goment.New()
	if err != nil {
		log.Println("init goment", err)
		return ""
	}
	return g.Format(format)
}

func onReady() {
	systray.SetTemplateIcon(iconM, icon.Data)
	systray.SetTooltip(links["appName"])

	github := systray.AddMenuItem("Github", "gozeon")
	twitter := systray.AddMenuItem("Twitter", "gozeonl")

	alert := systray.AddMenuItem("Alert", "系统通知")
	clipboardMenu := systray.AddMenuItem("Clipboard", "粘贴板")

	timeMenu := systray.AddMenuItem("Time", "时间")
	ymd := timeMenu.AddSubMenuItem("YYYY-MM-DD", "")
	hms := timeMenu.AddSubMenuItem("HH:mm:ss", "")
	ymdhms := timeMenu.AddSubMenuItem("YYYY-MM-DD HH:mm:ss", "")
	ymdhms1 := timeMenu.AddSubMenuItem("YYYYMMDDHHmmss", "")

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
			case <-ymd.ClickedCh:
				t := getTime("YYYY-MM-DD")
				err := clipboard.WriteAll(t)
				if err != nil {
					log.Println("clipboard wirte ymd", err)
				}
				log.Println("ymd", t)
			case <-hms.ClickedCh:
				t := getTime("HH:mm:ss")
				err := clipboard.WriteAll(t)
				if err != nil {
					log.Println("clipboard wirte hms", err)
				}
				log.Println("hms", t)
			case <-ymdhms.ClickedCh:
				t := getTime("YYYY-MM-DD HH:mm:ss")
				err := clipboard.WriteAll(t)
				if err != nil {
					log.Println("clipboard wirte ymdhms", err)
				}
				log.Println("ymdhms", t)
			case <-ymdhms1.ClickedCh:
				t := getTime("YYYYMMDDHHmmss")
				err := clipboard.WriteAll(t)
				if err != nil {
					log.Println("clipboard wirte ymdhms1", err)
				}
				log.Println("ymdhms1", t)
			}
		}
	}()
}

func onExit() {
	log.Println("onExit")
}
