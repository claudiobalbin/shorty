build:
	go build -o shorty main.go

local-environment:
	docker compose -f docker-compose-local.yml down
	docker compose -f docker-compose-local.yml stop
	docker compose -f docker-compose-local.yml build --force-rm
	docker compose -f docker-compose-local.yml up
