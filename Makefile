DIST = dist


.PHONY: chat
chat:
	go build -o chat/${DIST}/main chat/main.go
	chat/${DIST}/main
