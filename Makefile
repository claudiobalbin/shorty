build:
	go build -o shorty main.go

local-environment:
	docker compose -f docker-compose-local.yml down
	docker compose -f docker-compose-local.yml stop
	docker compose -f docker-compose-local.yml build --force-rm
	docker compose -f docker-compose-local.yml up

integration-tests:
	docker compose -f docker-compose-tests.yaml stop
	docker compose -f docker-compose-tests.yaml rm -f
	docker compose -f docker-compose-tests.yaml build
	docker compose -f docker-compose-tests.yaml up --exit-code-from app-test
