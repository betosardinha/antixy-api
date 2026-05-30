COMPOSE=docker compose -f docker-compose.dev.yml

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up -d db redis api

down:
	$(COMPOSE) down

migrate-up:
	$(COMPOSE) run --rm migrate goose up

migrate-down:
	$(COMPOSE) run --rm migrate goose down

migrate-status:
	$(COMPOSE) run --rm migrate goose status

logs:
	$(COMPOSE) logs -f

logs-%:
	$(COMPOSE) logs -f $*

tools:
	$(COMPOSE) up -d pgadmin

lint:
	$(COMPOSE) run --rm lint

test:
	$(COMPOSE) run --rm test
