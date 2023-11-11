run:
	@go run ./...

db_up:
	@docker-compose up -d

db_down:
	@docker-compose down --volumes && rm -rf ./data