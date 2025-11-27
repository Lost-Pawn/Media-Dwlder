package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"net/http"
	"time"
	env "github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	
)

func main() {
	err := env.Load() // loads .env
	if err != nil {
		log.Fatal("Error Loading .env: %v", err)
	}
	
	botToken := os.Getenv("BOT_TOKEN") // access token from .env
	if botToken == "" {
		fmt.Print("Bot Token Not Found.")
		return
	}
	config := tele.Settings{
		Token: botToken, // takes in bot token
		Poller: &tele.LongPoller{Timeout: 10*time.Second}, // reaches telegram servers and checks for messages every 10s before processing
		Verbose: true, // debugs itself and prints json on terminal
	}

	bot, err := tele.NewBot(config) // initializes
	if err != nil {
		log.Fatal("Error Creating Bot: %v", err)
	}
	log.Fatal("Bot is Running")
	
	bot.Handle(tele.OnText, func(c tele.Context) error {
		url := c.Text() // gets text
		chatID := c.Chat().ID // gets chatid & user id
		msg := &tele.Chat{ID: chatID} // stores chatid
	    bot.Send(msg, "Downloading your media...")

		result, err := Download(url)
		if err != nil {
			bot.Send(msg, "Failed to download your media.")
			return err
		}
		fileName := fmt.Sprintf("media_%d", time.Now().Unix())
		
		out, err := os.Create(fileName)
		if err != nil {
			bot.Send(msg, "Failed to create file.")
			return err
		}
		
		defer out.Close()
		// Downdloading via http 
		response, err := http.Get(result)
		if err != nil {
			bot.Send(msg, "Failed to save media file.")
			return err
		}
		
		defer response.Body.Close()
		_, err = io.Copy(out, response.Body)
		if err != nil {
			bot.Send(msg, "Failed to Downlaod.")
			return err
		}

		// Send actual media file
    	_, err = bot.Send(msg, &tele.Document{File: tele.FromDisk(fileName)})
    	if err != nil {
        	bot.Send(msg, "Failed to send media file.")
        	return err
    	}

		os.Remove(fileName)
		return nil
	}) 

	
}
