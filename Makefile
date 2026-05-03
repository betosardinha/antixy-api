compose=docker compose -f docker-compose.dev.yml

lint:
	$(compose) run --rm lint

migrate-up:
	$(compose) run --rm migrate goose up

migrate-down:
	$(compose) run --rm migrate goose down

migrate-status:
	$(compose) run --rm migrate goose status

up:
	$(compose) up -d --build db redis api

down:
	$(compose) down

logs:
	$(compose) logs -f

logs-api:
	$(compose) logs -f api

tools:
	$(compose) up -d pgadmin
