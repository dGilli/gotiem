build-tiem:
	@go build -o bin/tiem ./cmd/tiem/
build-ui:
	@go build -o bin/ui ./cmd/ui/

run-tiem: build-tiem
	@./bin/tiem
run-ui: build-tiem build-ui
	@./bin/ui
run: run-ui

clean:
	@rm -rf go.sum bin tmp
