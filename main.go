package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		data, err := ioutil.ReadFile("./replymessage.json")
		if err != nil {
			log.Fatal(err)
		}
		container, err := linebot.UnmarshalFlexMessageJSON(data)
		if err != nil {
			log.Fatal(err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("エンジニア寄せ鍋", container)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		// テンプレートをパース
		t := template.Must(template.ParseFiles("./personal_information.html"))
		// テンプレートを描画
		if err := t.ExecuteTemplate(w, "personal_information.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	})

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// func reply(replyToken string) error {
// 	container := &linebot.BubbleContainer{
// 		Type: linebot.FlexContainerTypeBubble,
// 		Body: &linebot.BoxComponent{
// 			Type:   linebot.FlexComponentTypeBox,
// 			Layout: linebot.FlexBoxLayoutTypeHorizontal,
// 			Contents: []linebot.FlexComponent{
// 				&linebot.TextComponent{
// 					Type: linebot.FlexComponentTypeText,
// 					Text: "Hello,",
// 				},
// 				&linebot.TextComponent{
// 					Type: linebot.FlexComponentTypeText,
// 					Text: "World!",
// 				},
// 			},
// 		},
// 	}
// 	if _, err := client.ReplyMessage(
// 		replyToken,
// 		linebot.NewFlexMessage("alt text", container),
// 	).Do(); err != nil {
// 		return err
// 	}
// 	return nil
// }
