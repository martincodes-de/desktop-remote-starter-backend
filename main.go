package main

import (
	"fmt"
	"github.com/martincodes-de/desktop-remote-starter-backend/src"
	"gopkg.in/ini.v1"
)

func main() {
	config, err := ini.Load("config.ini")

	if err != nil {
		fmt.Println("config.ini couldn't be loaded")
		return
	}

	isBrowserVisible, _ := config.Section("").Key("isBrowserVisible").Bool()
	fritzBoxUrl := config.Section("").Key("fritzBoxUrl").String()
	userName := config.Section("").Key("userName").String()
	userPassword := config.Section("").Key("userPassword").String()

	src.TurnOnComputer(isBrowserVisible, fritzBoxUrl, userName, userPassword, "DESKTOP-CCECHMS")
}
