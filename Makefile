build:
	go build -o ./bin/hotpocket -v ./

dev:
	go build -o ./bin/hotpocket -v ./
	./bin/hotpocket

clean:
	rm -r ./bin/