start-local:
	@docker-compose down || true
	@docker-compose up -d