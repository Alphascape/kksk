build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./app/kksk
	@go build -o docs/kksk ./app/kksk

run: build
	@cd docs && ./kksk local


build-github: build
	@cd docs && ./kksk github

github: build-github clean 

test:
	go test ./app/kksk
	GOARCH=wasm GOOS=js go test ./app/kksk

clean:
	@go clean ./...
	@-rm docs/kksk