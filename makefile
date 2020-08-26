SHELL := /bin/bash

api:
	docker build \
		-f zarf/compose/dockerfile-api \
		-t class-service-amd64:1.0 \
		--build-arg PACKAGE_NAME=service \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.

up:
	docker-compose -f zarf/compose/compose.yaml -f zarf/compose/compose-config.yaml up --detach --remove-orphans

down:
	docker-compose -f zarf/compose/compose.yaml down --remove-orphans

logs:
	docker-compose -f zarf/compose/compose.yaml logs -f

service:
	go run app/service/main.go

schema:
	go run app/admin/main.go schema

seed:
	go run app/admin/main.go seed

