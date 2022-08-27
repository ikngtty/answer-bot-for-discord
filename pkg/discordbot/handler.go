package discordbot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ikngtty/answer-bot-for-discord/pkg/chiebukuro"
)

func HandleMessageCreate(s *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}

	if !strings.HasSuffix(event.Content, "？") && !strings.HasSuffix(event.Content, "?") {
		return
	}

	question, err := chiebukuro.Sample()
	if err != nil {
		fmt.Println(err)
		return
	}

	message := fmt.Sprintf("%s\n<%s>", question.BestAnswer, question.URL)
	s.ChannelMessageSend(event.ChannelID, message)
}
