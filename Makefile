BINARY_NAME=githooks

build:
	go build -o $(BINARY_NAME) main.go

fmt:
	go fmt .

lint:
	golint .

test:
	go test .

clean:
	rm -f $(BINARY_NAME)
	rm -f hooks_config.json

run: build
	./$(BINARY_NAME)

push:
	git add .
	git commit -m "auto update"
	git push origin main
