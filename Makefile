install:
	go install github.com/google/wire/cmd/wire@latest

gen-wire:
	wire ./...

setup-db:
	bash ./tool/store/create_writer_tables.sh