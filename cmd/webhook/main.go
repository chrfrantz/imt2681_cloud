package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const theDiscordWebhook = "https://discordapp.com/api/webhooks/360514871512793108/asZ3STFa5of5sOzE1gNSi1MRe5Xgr3k--QthJeu5i_qwUGJe4C2twkr7RdmGAWZVEqQA"

/*
WebhookInfo represents the Discord webhook data entry.
*/
type WebhookInfo struct {
	Content string `json:"content"`
}

func sendDiscordLogEntry(what string) {
	info := WebhookInfo{}
	info.Content = what + "\n"
	raw, _ := json.Marshal(info)
	resp, err := http.Post(theDiscordWebhook, "application/json", bytes.NewBuffer(raw))
	if err != nil {
		fmt.Println(err)
		fmt.Println(ioutil.ReadAll(resp.Body))
	}
}

func main() {
	for {
		text := "Heroku timer test at: " + time.Now().String()
		delay := time.Minute * 30

		sendDiscordLogEntry(text)
		time.Sleep(delay)
	}
}
