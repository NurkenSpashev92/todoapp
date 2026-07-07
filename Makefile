.DEFAULT_GOAL := help

-include .env
export

DOCKER_COMP = docker-compose
DOCKER_EXEC = $(DOCKER_COMP) exec
APP = $(DOCKER_EXEC) todo

PROJECT_DIR=$(shell dirname $(realpath $(MAKEFILE_LIST)))

.PHONY: help env-up env-down migrate-create migrate-up migrate-down migrate-action


help: ## Help message
	@echo "Please choose a task:"
	@grep -hE '^[a-zA-Z_-]+:.*## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*## "}; \
	{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install project dependencies and prepare environment
	@echo "Installing Go dependencies..."
	@go mod download

	@echo "Starting PostgreSQL..."
	@$(DOCKER_COMP) up -d todo-postgres

	@echo "Running migrations..."
	@$(MAKE) migrate-up

	@echo "Project is ready."

run: ## Run application
	@$(DOCKER_COMP) up --build -d todo


env-up: ## Build postgres containers
	@$(DOCKER_COMP) up -d todo-postgres


env-down: ## Down postgres containers
	@$(DOCKER_COMP) down todo-postgres


migrate-create: ## Create migrations (use: make migrate-create seq=name)
	@if [ -z "$(seq)" ]; then \
		echo "Not found required param 'seq'"; \
		exit 1; \
	fi

	@$(DOCKER_COMP) run --rm todo-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq $(seq)


migrate-up: ## Run migrations
	@$(MAKE) migrate-action action=up


migrate-down: ## Rollback migrations
	@$(MAKE) migrate-action action=down


migrate-action:
	@$(DOCKER_COMP) run --rm todo-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todo-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		$(action)
