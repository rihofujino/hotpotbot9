package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/linebot" // ① SDKを追加
)

func main() {
	fmt.Println("server start ...")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// ② LINE bot instanceの作成
	// bot, err := linebot.New(
	// 	os.Getenv("CHANNEL_SECRET"),
	// 	os.Getenv("CHANNEL_TOKEN"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := mux.NewRouter()
	r.HandleFunc("/webhook", mainhandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Print(err)
	}
	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	// ③ LINE Messaging API用の Routing設定
	// router.POST("/webhook", func(c *gin.Context) {
	// 	events, err := bot.ParseRequest(c.Request)
	// 	if err != nil {
	// 		if err == linebot.ErrInvalidSignature {
	// 			log.Print(err)
	// 		}
	// 		return
	// 	}
	// 	for _, event := range events {
	// 		if event.Type == linebot.EventTypeMessage {
	// 			switch message := event.Message.(type) {
	// 			case *linebot.TextMessage:
	// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	// 					log.Print(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// })

	// router.Run(":" + port)
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("来た")
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Print(err)
	}
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Print(err)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
