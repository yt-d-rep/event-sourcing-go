install:
	go install github.com/google/wire/cmd/wire@latest

gen-wire:
	wire ./...