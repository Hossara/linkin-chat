# Variables
ROOT_DIR := ./build

NETWORK_NAME=linkin-chat-network

DEV_COMPOSE:= \
		-f $(ROOT_DIR)/nats/docker-compose.yaml \
		-f $(ROOT_DIR)/postgres/docker-compose.dev.yaml \

PROD_COMPOSE:= \
		-f $(ROOT_DIR)/nats/docker-compose.yaml \
		-f $(ROOT_DIR)/postgres/docker-compose.prod.yaml \
		-f $(ROOT_DIR)/server/docker-compose.yaml

# Functions
ensure-network:
	@echo "Ensuring project network is exists..."
	@if [ -z "$$(docker network ls --filter name=$(NETWORK_NAME) --format '{{.Name}}')" ]; then \
		echo "Network $(NETWORK_NAME) does not exist. Creating..."; \
		docker network create $(NETWORK_NAME); \
	fi

# CLI
create_release:
	goreleaser release --snapshot --clean

.PHONY: create_release

# Server Development
dev_up: ensure-network
	@docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up -d

dev_uplog: ensure-network
	@docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up

dev_down: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) down

dev_build: ensure-network
	@docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) build $(FLAGS) $(CONTAINER)

dev_logs: ensure-network
	@docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) logs -f

dev_rebuild: ensure-network
	@docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up --build

.PHONY: dev_up dev_uplog dev_down dev_build dev_logs dev_rebuild

# Server Deployment
deploy: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) up -d

deploy-down: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) down

deploy-logs: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) logs -f

deploy-log: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) logs -f $(CONTAINER)

deploy-build: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) build $(FLAGS) $(CONTAINER)

.PHONY: deploy deploy-down deploy-logs deploy-logs deploy-build