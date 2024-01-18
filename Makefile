all:
	go build -o bin/muis medius-server/cmd/muis
	go build -o bin/mas medius-server/cmd/mas
	go build -o bin/mls medius-server/cmd/mls