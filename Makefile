build-app:
	@go build -o bin/tiem ./cmd/tiem/

run: build-app
	@./bin/tiem

clean:
	@rm -rf bin
