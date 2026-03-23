project := ./docker-compose.yml
dcmd := docker compose -f $(project)
dexec := $(dcmd) exec -ti
dcrun := $(dcmd) run


default: help

.PHONY: build
# Build containers or specify with container=<container>.
build:
	$(dcmd) build $(container)

.PHONY: up
# Start containers or specify with container=<container>.
up:
	$(dcmd) up -d $(container)

.PHONY: down
# Stop and remove containers or specify with container=<container>.
down:
	$(dcmd) down $(container)

.PHONY: stop
# Stop containers without removing them or specify with container=<container>.
stop:
	$(dcmd) stop $(container)

.PHONY: logs
# Show logs from all running containers or specify with container=<container>.
logs:
	$(dcmd) logs -f --tail 10	 $(container)

.PHONY: status
# Show running containers.
status:
	$(dcmd) ps

.PHONY: enter
# Open a shell into a container with container=<container>.
enter:
	$(dexec) $(container) /bin/bash

.PHONY: exec
# Execute a command from a container with container=<container> cmd=<cmd>.
exec:
	$(dexec) $(container) $(cmd)

#run:
#	$(dcrun) $(container) $(cmd)

.PHONY: clean
# Stop and remove containers along with their named volumes and networks.
clean:
	$(dcmd) down -v


.PHONY: nuke
# Try to cleanup full docker environment.
nuke: nuke_stop_running nuke_rm_all nuke_rm_images nuke_rm_volumes nuke_prune

.PHONY: nuke_stop_running
nuke_stop_running:
	@echo "- Stopping containers"
	$(eval output = $(shell docker ps -q))
	@if [ ! -z "$(output)" ]; then \
		docker stop $(output); \
	fi

.PHONY: nuke_rm_all
nuke_rm_all:
	@echo "- Removing containers"
	$(eval output = $(shell docker ps -a -q))
	@if [ ! -z "$(output)" ]; then \
		docker rm $(output); \
	fi

.PHONY: nuke_rm_images
nuke_rm_images:
	@echo "- Removing images"
	$(eval output = $(shell docker image ls -q))
	@if [ ! -z "$(output)" ]; then \
		docker image rm $(output); \
	fi

.PHONY: nuke_rm_volumes
nuke_rm_volumes:
	@echo "- Removing volumes"
	$(eval output = $(shell docker volume ls -q))
	@if [ ! -z "$(output)" ]; then \
		docker volume rm $(output); \
	fi

.PHONY: nuke_rm_stopped_containers
nuke_rm_stopped_containers:
	@echo "- Removing stopped containers"
	docker container prune


.PHONY: nuke_prune
nuke_prune:
	docker system prune -a

.PHONY: help
# Show this help.
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

#.PHONY: api-docs
#api-docs:
#	npx open-swagger-ui ../authorization-gateway/swagger.json --open

