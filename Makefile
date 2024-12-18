DOCKER_COMPOSE_FILE := deployments/compose.yml
PROJECT_NAME := currency_exchange_rate_service

# Deploy the service
deploy:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) up

# Stop the service
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) down

# Restart the service
restart:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) down
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) up

# Stop and clean up the service (remove containers, networks, volumes, etc.)
stop_and_cleanup:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p $(PROJECT_NAME) down --volumes --remove-orphans

.PHONY: deploy stop restart stop_and_cleanup
