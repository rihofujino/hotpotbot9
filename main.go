package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	ChannelSecret := os.Getenv("CHANNEL_SECRET")
	ChannelAccessToken := os.Getenv("CHANNEL_TOKEN")

	router.POST("/webhook", func(c *gin.Context) {
		client := &http.Client{Timeout: time.Duration(15 * time.Second)}
		fmt.Println(client) //&{<nil> <nil> <nil> 15s}

		bot, err := linebot.New(ChannelSecret, ChannelAccessToken, linebot.WithHTTPClient(client))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(bot) //&{6aa37b7773a27d966174457ccb5c284e ix44e08i7lia3MIrspbXvrrSHfTxlpPMgSAoyggevENTdxnatN0RSgowwsk/KZgtBSCpoB9osesTOo30NuzFxxgnIWCmUCtYrCU8AhARnMfx38MumaPD1RnfzvCKYyxhhLCaMF58vIlEJhzKYq7piQdB04t89/1O/w1cDnyilFU= 0xc0001f6300 0xc0001b6e40}
		received, err := bot.ParseRequest(c.Request)
		fmt.Println(received) //[0xc0001e65a0]

		for _, event := range received {
			fmt.Println(event) //&{9f5422904b2d4b70ac3a7305ed334b42 message 2019-03-30 12:55:39.218 +0000 UTC 0xc0001bc600 0xc0001b7110 <nil> <nil> <nil> <nil> <nil> <nil> []}
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					source := event.Source
					if source.Type == linebot.EventSourceTypeRoom {
						if resMessage := getResMessage(message.Text); resMessage != "" {
							postMessage := linebot.NewTextMessage(resMessage)
							if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
								log.Print(err)
							}
						}
					}
				}
			}
		}
	})

	router.Run(":" + port)
}

func getResMessage(reqMessage string) (message string) {
	resMessages := [3]string{"わかるわかる", "それで？それで？", "からの〜？"}

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(5) == 0 {
		if math := rand.Intn(4); math != 3 {
			message = resMessages[math]
		} else {
			message = reqMessage + "じゃねーよw"
		}
	}
	return
}
