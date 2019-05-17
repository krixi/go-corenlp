
retool:
	which retool > /dev/null || go get -u github.com/twitchtv/retool
	go install github.com/twitchtv/retool
	retool sync

fmt: retool
	retool do goimports -w .

test:
	docker-compose up -d
	go test -race ./...
.PHONY: test


