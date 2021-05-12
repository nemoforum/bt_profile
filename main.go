//go:generate go run gen/main.go

package main

import (
	"flag"
	"os/exec"

	"github.com/getlantern/systray"
)

var sink string
var codecm string
var codecv string
var showQuit bool

func main() {
	s := flag.String("sink", "1", "headset's pulseaudio sink")
	q := flag.Bool("quit", false, "show the quit item")
	cm := flag.String("a2dp-codec", "", "custom codec e.g. aac, ldac")
	cv := flag.String("hsp-hfp-codec", "", "custom codec e.g. aac, ldac")
	flag.Parse()
	sink = *s
	showQuit = *q
	if *cm != "" {
		codecm = "-" + *cm
	}
	if *cv != "" {
		codecv = "-" + *cv
	}
	systray.Run(onready, nil)
}

func onready() {
	systray.SetIcon(icon)
	a2dp := systray.AddMenuItem("A2DP", "Switch to A2DP mode")
	hsphfp := systray.AddMenuItem("HSP/HFP", "Switch to HSP/HFP mode")
	quit := systray.AddMenuItem("Quit", "Quit the app")
	if !showQuit {
		quit.Hide()
	}
	for {
		select {
		case <-quit.ClickedCh:
			systray.Quit()
		case <-a2dp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "a2dp-sink"+codecm).Run()
		case <-hsphfp.ClickedCh:
			exec.Command("pactl", "set-card-profile", sink, "headset-head-unit"+codecv).Run()
		}
	}
}
