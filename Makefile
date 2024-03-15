bin/%:
	@mkdir -p $(@D)
	@go build -o $@ ./cmd/$(notdir $@)

run: bin/$(filter-out run,$(MAKECMDGOALS))
	@mkdir -p ./tmp
	@./bin/$(filter-out $@,$(MAKECMDGOALS))

clean:
	@rm -rf bin tmp

%:
	@:
# ref: https://stackoverflow.com/questions/6273608/how-to-pass-argument-to-makefile-from-command-line

