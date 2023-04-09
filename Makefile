build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: build
	./contextual-state-chart-editor-go

testTT:
	go test -C ContextualStateChart/TrieTree