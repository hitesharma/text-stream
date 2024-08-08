VERSION ?= latest
REPO ?= hitesharma
SERVER_IMG = $(REPO)/text-stream-server:$(VERSION)
CLIENT_IMG = $(REPO)/text-stream-client:$(VERSION)

run-ws-server:
	go run cmd/ws-server/main.go

run-ws-client:
	go run cmd/ws-client/main.go

build-ws-server: go-format go-vet
	docker build -t ${SERVER_IMG} -f build/ws-server/dockerfile .

build-ws-client: go-format go-vet
	docker build -t ${CLIENT_IMG} -f build/ws-client/dockerfile .

go-format:
	go fmt ./...

go-vet:
	go vet ./...
