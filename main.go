package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"time"
	"ulxng/forwarder-bot/db"
	"ulxng/forwarder-bot/repository"

	"github.com/jessevdk/go-flags"
	tele "gopkg.in/telebot.v4"
	_ "modernc.org/sqlite"
)

type App struct {
	configRepository repository.ForwardConfigRepository
}

type options struct {
	BotToken        string `long:"token" env:"BOT_TOKEN" required:"true" description:"telegram bot token"`
	DBConnectionDSN string `long:"dsn" env:"DB_DSN" required:"true" description:"database connection string"`
}

func main() {
	var opts options
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	if _, err := p.Parse(); err != nil {
		log.Fatal(err)
	}

	log.Println("bot started")

	if err := run(opts); err != nil {
		log.Printf("run: %s", err)
	}

	log.Println("bot stopped")
}

func run(opts options) error {
	queries, err := db.CreateConnection(opts.DBConnectionDSN)
	if err != nil {
		return fmt.Errorf("createConnection: %w", err)
	}
	a := &App{
		configRepository: repository.NewForwardConfigRepository(queries),
	}
	pref := tele.Settings{
		Token:  opts.BotToken,
		Poller: &tele.LongPoller{Timeout: time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		return fmt.Errorf("tele.NewBot: %w", err)
	}

	bot.Handle("/start", a.help)
	bot.Handle("/ping", a.ping)
	bot.Handle("/init", a.init)
	bot.Handle(tele.OnBusinessMessage, a.handleReceived)
	bot.Handle(tele.OnEditedBusinessMessage, a.handleEdited)

	bot.Start()
	return nil
}

func (a *App) help(c tele.Context) error {
	return c.Send("How to use the bot:\n\n" +
		"1. Go to Telegram Settings -> Telegram Business -> ChatBots\n" +
		"2. Add this bot \n" +
		"3. Select the contacts you want to manage\n" +
		"Done!\n" +
		"\n" +
		"❕ Telegram Premium required.",
	)
}

func (a *App) ping(c tele.Context) error {
	return c.Send("pong")
}

func (a *App) handleEdited(c tele.Context) error {
	return fmt.Errorf("handleEdited: not implemented")
}

func (a *App) handleReceived(c tele.Context) error {
	receivedMsg := c.Update().BusinessMessage
	if receivedMsg.Chat.ID != receivedMsg.Sender.ID {
		//skip owner messages
		return nil
	}
	var m any
	if receivedMsg.Media() != nil {
		m = receivedMsg.Media()
	} else {
		m = receivedMsg.Text
	}
	inboxChat, err := a.extractInboxChatID(receivedMsg.BusinessConnectionID, c.Bot())
	if err != nil {
		return fmt.Errorf("extractInboxChatID: %w", err)
	}
	sent, err := c.Bot().Send(inboxChat, m)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}

	if _, err := c.Bot().Reply(sent, fmt.Sprintf(
		"from %s %s @%s",
		receivedMsg.Sender.FirstName,
		receivedMsg.Sender.LastName,
		receivedMsg.Sender.Username, // может быть пустым
	)); err != nil {
		return fmt.Errorf("reply: %w", err)
	}
	return nil
}

func (a *App) init(c tele.Context) error {
	userID := c.Sender().ID
	conf := repository.ForwardConfig{
		UserID: userID,
		ChatID: c.Chat().ID,
	}
	ctx := context.Background()
	if err := a.configRepository.Save(ctx, conf); err != nil {
		return fmt.Errorf("save: %w", err)
	}

	return c.Send(fmt.Sprintf("this chat selected to forward messages from %s", c.Sender().Username))
}

func (a *App) extractInboxChatID(businessConnectionID string, bot tele.API) (tele.Recipient, error) {
	bc, err := bot.BusinessConnection(businessConnectionID)
	if err != nil {
		return nil, fmt.Errorf("businessConnection: %w", err)
	}
	if bc == nil {
		return nil, fmt.Errorf("businessConnection not found")
	}
	ctx := context.Background()
	chatID, err := a.configRepository.FindChatByUserID(ctx, bc.UserChatID)
	if err != nil {
		return nil, fmt.Errorf("findByUser: %w", err)
	}
	if chatID == 0 {
		return bc.Sender, nil
	}
	return tele.ChatID(chatID), nil
}
