#
# Just in case
#
build:
	@air build	

clean:
	@rm -rf ./bin/*

test:
	@go test

# Cli that application relies of
install-cli:
	@go install github.com/air-verse/air@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
