lint:
	golint ./...

imports:
	goimports -w .

test:
	go test ./...
