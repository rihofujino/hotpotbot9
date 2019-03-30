package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// ① SDKを追加
)

func main() {
	fmt.Println("server start ...")
	port := os.Getenv("PORT")
	fmt.Println(port)

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)

	// r := mux.NewRouter()
	// r.HandleFunc("/", testpage)
	// r.HandleFunc("/webhook", mainhandler).Methods("POST")

	// mux := http.NewServeMux()
	// fmt.Println(mux)
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("来た")
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })

	// mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("来た")
	// 	bot, err := linebot.New(
	// 		os.Getenv("CHANNEL_SECRET"),
	// 		os.Getenv("CHANNEL_TOKEN"),
	// 	)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	events, err := bot.ParseRequest(r)
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

	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Print(err)
	// }
}

// //MainPage ...
// func testpage(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hi"))
// }

// func mainhandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("来た")
// 	bot, err := linebot.New(
// 		os.Getenv("CHANNEL_SECRET"),
// 		os.Getenv("CHANNEL_TOKEN"),
// 	)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	events, err := bot.ParseRequest(r)
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
// }
