package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Main ...
func Main(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	data, err := ioutil.ReadFile("messages/replymessage.json")
	if err != nil {
		log.Fatal(err)
	}
	container, err := linebot.UnmarshalFlexMessageJSON(data)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			fmt.Println("debug1")
			fmt.Println(event.ReplyToken)
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("エンジニア寄せ鍋", container)).Do(); err != nil {
				fmt.Println("debug2")
				log.Fatal(err)
			}
		}
	}
}
