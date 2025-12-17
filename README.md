# Forwarder Bot
A Telegram bot that forwards messages from a private business account to a selected chat.

Works only with Telegram Premium accounts.
## Run
```bash
go run main.go --token=[BOT_TOKEN]
```

## Usage
By default, the bot forwards messages to the private chat with the bot itself.

To use a group as an inbox:
1.	Add the bot to the group.
2.	Run the `/init` command in that group.

The group will be assigned as the inbox for your account.

One user can have only one inbox chat at a time.
