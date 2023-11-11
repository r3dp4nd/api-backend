run:
	@go run ./...

db_up:
	@rm -rf data/ && docker-compose up -d

db_down:
	@docker-compose down --volumes