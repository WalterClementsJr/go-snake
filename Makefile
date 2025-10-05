binary_name = snake

.PHONY: tidy clean install build run

tidy:
	go mod tidy -v
	go fmt ./...

clean:
	go clean
	/bin/rm -rf ./bin/

install:
	@sudo apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev
	@go install

build:
	go build -o ./bin/${binary_name} .

run: build
	./bin/${binary_name}

test:
	go test -v -race -coverprofile=coverage.out -coverpkg=./... ./...

