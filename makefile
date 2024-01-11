# test を実行
.PHONY: test
test: 
	go list -f '{{.Dir}}/...' -m | WORKSPACE_DIR=$(shell pwd) xargs go test -cover -v

.PHONY: lint
lint: 
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/unparam@latest
	golangci-lint run
	unparam -exported ./...

# 脆弱性診断を実行
.PHONY: vulncheck
vulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go list -f '{{.Dir}}/...' -m | xargs govulncheck


# wire を使って DI するためのコードを生成する
.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd internal/injector; wire

# entのgenerateを行う
.PHONY: ent
ent:
    GOWORK=off go generate ./tool/ent/

.PHONY: install
install:
	go mod tidy