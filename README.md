# Forwarder Bot
Forwards messages from private business account to selected chat.

Needs Telegram Premium.

## Run
Forward to concrete chat:
```bash
go run main.go --token=[BOT_TOKEN] --chat=[INBOX_CHAT_ID]
```

Forward to private chat (bot + user)
```bash
go run main.go --token=[BOT_TOKEN] --to-owner=1
```