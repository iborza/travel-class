SHELL := /bin/bash

up:
	docker-compose -f zarf/compose/compose.yaml -f zarf/compose/compose-config.yaml up --detach --remove-orphans

down:
	docker-compose -f zarf/compose/compose.yaml down --remove-orphans

schema:
	go run app/admin/main.go schema

seed: schema
	go run app/admin/main.go seed

