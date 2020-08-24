SHELL := /bin/bash

schema:
	go run app/admin/main.go schema

seed: schema
	go run app/admin/main.go seed

