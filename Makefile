include scripts/help.mk

.PHONY: build
build: ##@docker build the docker image
	docker-compose up --build -d

.PHONY: stop
stop: ##@docker stop the docker image
	docker-compose stop

.PHONY: logs
logs: ##@docker show the last 100 lines
	docker-compose logs -f --tail=100

.PHONY: clean
clean: ##@docker destroy docker image and clean the volumes
	docker-compose down --remove-orphans --volumes

.PHONY: bash
bash: ##@docker open the bash in the docker image
	docker-compose run --rm db /bin/bash

.PHONY: test
test:
	go test -tags testing ./...
