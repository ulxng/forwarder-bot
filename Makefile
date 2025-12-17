include .env

format:
	gofmt -s -w . && \
	goimports -w .

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./release/bot .

deploy:
	ssh ${SERVER} systemctl --user stop ${SERVICE_NAME} || true
	ssh ${SERVER} "mkdir -p ${DEPLOY_DIR} && mkdir -p ~/.config/systemd/user/"
	rsync -rzh --progress .env ./release/bot ${SERVER}:${DEPLOY_DIR}
	rsync -z release/systemd/bot.service ${SERVER}:~/.config/systemd/user/${SERVICE_NAME}
	ssh ${SERVER} "systemctl --user daemon-reload && systemctl --user restart ${SERVICE_NAME}"

generate:
	sqlc generate --file db/sqlc/sqlc.yaml