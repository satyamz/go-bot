package main

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catgif"
	// _ "github.com/go-chat-bot/plugins/chucknorris"
	// Import all the commands you wish to use
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}

func hello(command *bot.Cmd) (msg string, err error) {
	msg = fmt.Sprintf("Hello %s", command.User.RealName)
	return
}

func greet(command *bot.Cmd) (bot.CmdResult, error) {
	commandResult := &bot.CmdResult{}
	commandResult.Channel = "general"
	commandResult.Message = "Hey! Go bot here.."
	return *commandResult, nil
}

func init() {

	bot.RegisterCommandV2(
		"greet",
		"Sends a greeting to channel ",
		"sub command",
		greet)

	bot.RegisterCommand(
		"hello",
		"Sends a 'Hello' message to you on the channel.",
		"",
		hello)
}
