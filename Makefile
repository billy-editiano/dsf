build:
	go build -o bin/$(app) cmd/$(app)/main.go

build-all:
	make build app=server
	echo "server build done\n"
	make build app=fetcher
	echo "fetcher build done\n"
	make build app=status
	echo "status build done"

install:
	make build app=$(app)
	sudo cp bin/$(app) /usr/local/bin/$(app)


