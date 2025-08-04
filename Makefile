.PHONY: up
up:
	@echo "docker compose up"
	@docker compose -f docker-compose.yml up -d

.PHONY: down
down:
	@echo "docker compose down"
	@docker compose -f docker-compose.yml down


