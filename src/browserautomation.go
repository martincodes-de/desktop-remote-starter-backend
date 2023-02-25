package src

import (
	"github.com/go-rod/rod"
	browserLauncher "github.com/go-rod/rod/lib/launcher"
	"strings"
	"time"
)

func TurnOnComputer(isBrowserVisible bool, fritzBoxUrl, userName, userPassword, machineName string) {

	launcher := browserLauncher.New().Headless(isBrowserVisible).MustLaunch()
	browser := rod.New().ControlURL(launcher).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(fritzBoxUrl)

	page.MustElement("input#uiViewUser").MustInput(userName)
	page.MustElement("input#uiPass").MustInput(userPassword)
	page.MustElement("#submitLoginBtn").MustClick()

	time.Sleep(time.Second * 5)

	page.MustElement("a#lan").MustClick()
	page.MustElement("a#net").MustClick()

	time.Sleep(time.Second * 15)

	var devices, _ = page.Elements("div .network .list")
	var deviceToStart *rod.Element = nil

	for _, device := range devices {
		if strings.Contains(device.MustText(), machineName) {
			deviceToStart = device
		}
	}

	editButton, _ := deviceToStart.Element("[name=edit]")
	editButton.MustClick()

	time.Sleep((time.Second) * 5)

	page.MustElement("[name=btn_wake]").MustClick()

	time.Sleep(time.Second * 15)
}
