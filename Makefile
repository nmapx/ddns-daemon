include .env
export

include Makefile.common

COMPOSE_BIN := docker compose
COMPOSE_FILES := -f compose.yml

MAKEFILE_APP := Makefile.app

all: | build up

.PHONY: build
build:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} build --pull --force-rm --no-cache

.PHONY: up
up: | docker-up

.PHONY: down
down: | docker-down

.PHONY: restart
restart: | docker-restart

.PHONY: docker-up
docker-up:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} up -d --force-recreate

.PHONY: docker-down
docker-down:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} down -t 3

.PHONY: docker-restart
docker-restart:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} restart -t 1

.PHONY: sh
sh:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} exec app sh

.PHONY: bash
bash:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} exec app bash

.PHONY: app-daemon
app-daemon:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} exec -T app make -f ${MAKEFILE_APP} daemon

.PHONY: app-build
app-build:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} exec -T app make -f ${MAKEFILE_APP} build

.PHONY: app-test
app-test:
	${COMPOSE_BIN} -p ${PROJECT_NAME} ${COMPOSE_FILES} exec -T app make -f ${MAKEFILE_APP} test
