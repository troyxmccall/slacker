package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/troyxmccall/slacker"
	"github.com/slack-go/slack"
)

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	definition := &slacker.CommandDefinition{
		Description: "Upload a sentence!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			sentence := request.Param("sentence")
			apiClient := botCtx.APIClient()
			event := botCtx.Event()

			if event.ChannelID != "" {
				apiClient.PostMessage(event.ChannelID, slack.MsgOptionText("Uploading file ...", false))
				_, err := apiClient.UploadFile(slack.FileUploadParameters{Content: sentence, Channels: []string{event.ChannelID}})
				if err != nil {
					fmt.Printf("Error encountered when uploading file: %+v\n", err)
				}
			}
		},
	}

	bot.Command("upload <sentence>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
