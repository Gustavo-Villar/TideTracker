# Makefile for the TideTracker project

.PHONY: build test tests migrations_dev migrations_prod up_dev_db up_prod_system deploy

POSTGRES_DEV_URL=postgres://postgres:postgres@localhost:5432/tidetracker?sslmode=disable
POSTGRES_PROD_URL=postgres://postgres:postgres@localhost:5432/tidetracker?sslmode=disable

# Build the TideTracker application
build:
	@go build && echo "✨Build Sucessfull✨"

# Run tests
test tests:
	@go test ./...

# Apply migrations in the development environment
migrations_dev:
	@goose -dir sql/schema postgres "$(POSTGRES_DEV_URL)" up && echo "✨Dev Migrations Sucessfull✨"

# Apply migrations in the production environment
migrations_prod:
	@echo "⚠️ WARNING: You are about to apply migrations. Are you sure? [y/N]: " && read ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		goose -dir sql/schema postgres "$(POSTGRES_URL)" up && echo "✨PRODUCTION Migrations Sucessfull✨"; \
	else \
		echo "Operation aborted. ❌"; \
	fi

# Start the development database
up_dev_db:
	@docker-compose -f docker-compose-db-only.yml up -d

# Start the entire system in production mode
up_prod_system deploy:
	@echo "⚠️ You are about to deploy in PRODUCTION mode. Are you sure? [y/N]: " && read ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		docker-compose up -d --build; \
	else \
		echo "Operation aborted. ❌"; \
	fi