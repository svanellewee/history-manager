all:
	go mod vendor
	go build .

clean:
	rm history-manager


