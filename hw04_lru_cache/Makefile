
lint:
	go fmt ./...
	goimports -w .
	golangci-lint run --enable-all ./...

test:
	go test -v ./...
	go test -v -count=1 -race -gcflags=-l -timeout=30s

push:
	git push --set-upstream origin hw04_lru_cache
