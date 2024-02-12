.PHONY: run-local
run-local:
	go run ./cmd/main.go -v

.PHONY: run-docker-local
run-docker-dev:
	docker-compose -f ./scripts/docker/local/docker-compose.yml up --build  --remove-orphans

.PHONY: run-docker-local
run-requirements-local:
	docker-compose -f ./scripts/docker/local/docker-compose.requirements.yml up --build  --remove-orphans

.PHONY: generate-rest-docs
generate-rest-docs:
	swag init --parseDependency --parseInternal --parseDepth 2 -g .\cmd\main.go