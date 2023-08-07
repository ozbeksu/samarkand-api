hello:
	@echo "hello from bubble-bird"

build:
	@go build -o bin/api

run: clean build
	@clear
	@./bin/api

seed:
	@clear
	@go run ./scripts

sch:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

gen:
	go generate ./ent

clean:
	@rm -rf ./bin

re-seed: gen seed

re-run: gen seed run

test-all:
	@go test -v ./...

count = 1
test:
	@go test -v ./ -run $(name) -count $(count)

compile:
	echo "compiling for every os and platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go

docker:
	@docker-compose up --build
