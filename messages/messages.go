package messages

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func Greeting() string {
	return ""
}

func Help() string {
	return "How to use the bot:\n\n" +
		"1. Go to Telegram Settings -> Telegram Business -> ChatBots\n" +
		"2. Add this bot \n" +
		"3. Select the contacts you want to manage\n" +
		"Done!\n" +
		"\n" +
		"❕ Telegram Premium required."
}

func Confirm(sender string) string {
	return fmt.Sprintf("this chat selected to forward messages from %s", sender)
}

func Signature(message tele.Message) string {
	return fmt.Sprintf(
		"from %s %s @%s",
		message.Sender.FirstName,
		message.Sender.LastName,
		message.Sender.Username, // can be empty
	)
}
