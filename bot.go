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

		result, err := Download(url) // stores url in result
		if err != nil {
			bot.Send(msg, "Failed to download your media.")
			return err
		}
		fileName := fmt.Sprintf("media_%d", time.Now().Unix()) // formats a path file
		
		out, err := os.Create(fileName) // creates
		if err != nil {
			bot.Send(msg, "Failed to create file.")
			return err
		}
		
		defer out.Close() // on hold, closes
		// Downdloading via http 
		response, err := http.Get(result) // downloads result from url
		if err != nil {
			bot.Send(msg, "Failed to save media file.")
			return err
		}
		
		defer response.Body.Close() // holds body = result, closes
		_, err = io.Copy(out, response.Body) // keeps the result inside a file
		if err != nil {
			bot.Send(msg, "Failed to Downlaod.")
			return err
		}

		// Send actual media file
    	_, err = bot.Send(msg, &tele.Document{File: tele.FromDisk(fileName)}) // sends file
    	if err != nil {
        	bot.Send(msg, "Failed to send media file.")
        	return err
    	}

		os.Remove(fileName) // removes file
		return nil
	}) 

bot.Handle(tele.OnQuery, func(c tele.Context) error {
    query := c.Query().Text // @bot <query>

    if query == "" {
        return c.Answer(&tele.QueryResponse{
            Results:  []tele.Result{},
            CacheTime: 0,
        })
    }

    directURL, err := Download(query) // fetches
    if err != nil {
        return c.Answer(&tele.QueryResponse{
            Results:  []tele.Result{},
            CacheTime: 0,
        })
    }

    // ---- FALLBACK TEXT OPTION ----
    article := &tele.ArticleResult{
        Title:       "Direct Download Link",
        Description: "Tap to copy the original URL",
        Text:        directURL,
        HideURL:     true,
    }
    article.SetResultID("4")

    // FIX: results must be a tele.Results (slice of tele.Result)
    var results tele.Results
    results = append(results, article) // appends article format in result = NONE

    return c.Answer(&tele.QueryResponse{
        Results:  results,
        CacheTime: 0,
    })
})
}
		