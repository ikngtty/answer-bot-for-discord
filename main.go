package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ikngtty/answer-bot-for-discord/pkg/discordbot"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("env var BOT_TOKEN required")
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session:", err)
		os.Exit(1)
	}

	dg.AddHandler(discordbot.HandleReady)
	dg.AddHandler(discordbot.HandleMessageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		os.Exit(1)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	fmt.Println("Closed successfully.")
}
