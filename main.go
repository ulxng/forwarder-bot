package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"

	tele "gopkg.in/telebot.v4"
)

type App struct {
	InboxChatID tele.ChatID
}

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
	a := &App{
		InboxChatID: tele.ChatID(opts.InboxChatID),
	}
	pref := tele.Settings{
		Token:  opts.BotToken,
		Poller: &tele.LongPoller{Timeout: time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		return fmt.Errorf("tele.NewBot: %w", err)
	}

	bot.Handle("/start", a.ping)
	bot.Handle(tele.OnBusinessMessage, a.handleReceived)
	bot.Handle(tele.OnEditedBusinessMessage, a.handleEdited)

	bot.Start()
	return nil
}

func (a *App) ping(c tele.Context) error {
	return c.Send("pong!")
}

func (a *App) handleEdited(c tele.Context) error {
	return fmt.Errorf("handleEdited: not implemented")
}

func (a *App) handleReceived(c tele.Context) error {
	receivedMsg := c.Update().BusinessMessage
	var m any
	if receivedMsg.Media() != nil {
		m = receivedMsg.Media()
	} else {
		m = receivedMsg.Text
	}
	sent, err := c.Bot().Send(a.InboxChatID, m)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}

	if _, err := c.Bot().Reply(sent, fmt.Sprintf(
		"from\n%s %s @%s",
		receivedMsg.Sender.FirstName,
		receivedMsg.Sender.LastName,
		receivedMsg.Sender.Username, // может быть пустым
	)); err != nil {
		return fmt.Errorf("reply: %w", err)
	}
	return nil
}
