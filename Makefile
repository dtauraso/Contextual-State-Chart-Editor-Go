build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

buildWindows:
	set GOOS=js ; set GOARCH=wasm ; go build -o web/app.wasm
	go build

run: build
	./contextual-state-chart-editor-go

runWindows: buildWindows
	./contextual-state-chart-editor-go
testTT:
	go test -C ContextualStateChart/TrieTree