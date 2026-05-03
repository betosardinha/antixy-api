COMPOSE=docker compose -f docker-compose.dev.yml

lint:
	$(COMPOSE) run --rm lint

migrate-up:
	$(COMPOSE) run --rm migrate goose up

migrate-down:
	$(COMPOSE) run --rm migrate goose down

migrate-status:
	$(COMPOSE) run --rm migrate goose status

up:
	$(COMPOSE) up -d --build db redis api

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f

logs-%:
	$(COMPOSE) logs -f $*

tools:
	$(COMPOSE) up -d pgadmin

test:
	rm -f coverage.out
	$(COMPOSE) run --rm test

coverage:
	$(COMPOSE) run --rm test go tool cover -func=coverage.out

coverage-html:
	$(COMPOSE) run --rm test go tool cover -html=coverage.out -o coverage.html
	open coverage.html