.PHONY: install 

install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/tomwright/dasel/v2/cmd/dasel@master

.PHONY:  migrate

migrate:
	sh ./goose.sh up

.PHONY: delete

delete:
	sh ./goose.sh down

.PHONY: info

info:
	sh ./goose.sh status

