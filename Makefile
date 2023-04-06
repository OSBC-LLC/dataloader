build:
	@echo "Building go-rest-service"
	@go build -o bin/dataload -v .
	@echo "done."

test:
	@echo "running tests..."
	@go test -v ./tests -coverprofile=./coverage.out -coverpkg ./...

report:
	@echo "showing test coverage report..."
	@go tool cover -html=coverage.out

db:
	@echo "building docker images"
	@docker compose build
	@echo done.

du:
	@echo "starting docker images"
	@docker compose up

da: db du
