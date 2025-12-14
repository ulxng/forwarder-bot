package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"

	tele "gopkg.in/telebot.v4"
)

type options struct {
	BotToken    string `long:"token" env:"BOT_TOKEN" required:"true" description:"telegram bot token"`
	InboxChatID int    `long:"chat_id" env:"INBOX_CHAT_ID" required:"true" description:"chat to send message to"`
}

func main() {
	var opts options
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}

	log.Println("bot started")

	if err := run(opts); err != nil {
		log.Printf("run: %s", err)
	}

	log.Println("bot stopped")
}

func run(opts options) error {
	pref := tele.Settings{
		Token:  opts.BotToken,
		Poller: &tele.LongPoller{Timeout: time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		return fmt.Errorf("tele.NewBot: %w", err)
	}

	bot.Handle("/start", handle)

	bot.Start()
	return nil
}

func handle(c tele.Context) error {
	return c.Send("Hello!")
}
