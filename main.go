package main

import (
	"fmt"
	"github.com/martincodes-de/desktop-remote-starter-backend/src"
	"gopkg.in/ini.v1"
	"io"
	"net/http"
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
	httpApiKey := config.Section("").Key("httpApiKey").String()
	httpPort := config.Section("").Key("httpPort").String()

	http.HandleFunc("/check", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, "OK - ONLINE")
		return
	})

	http.HandleFunc("/turn-on-computer", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var queryParameter = request.URL.Query()

		if queryParameter.Get("apikey") != httpApiKey {
			writer.WriteHeader(http.StatusForbidden)
			io.WriteString(writer, "No/wrong apikey provided")
			return
		}

		var machineName = queryParameter.Get("machineName")
		if machineName == "" {
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "No machineName provided")
			return
		}

		go src.TurnOnComputer(isBrowserVisible, fritzBoxUrl, userName, userPassword, machineName)

		writer.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Server started with port " + httpPort)
	http.ListenAndServe(":"+httpPort, nil)
}
