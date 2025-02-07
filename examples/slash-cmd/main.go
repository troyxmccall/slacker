package main

import (
	"context"
	"log"
	"os"

	"github.com/troyxmccall/slacker"
)

// Implements a simple slash command. Assumes you have the slash command
// `/ping` defined for your app.

func main() {
	bot := slacker.NewClient(
		os.Getenv("SLACK_BOT_TOKEN"),
		os.Getenv("SLACK_APP_TOKEN"),
		slacker.WithDebug(true),
	)

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
		HideHelp: true,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
