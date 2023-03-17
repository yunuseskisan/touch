.PHONY: test

SERVICE = service.rpc.users
CI_COMMIT_SHA ?= latest

generate:
	cd ../ && bin/generate_go_proto $(SERVICE)

build:
	cd ../ && bin/build_service $(SERVICE) $(CI_COMMIT_SHA)

test:
	cd test &&  docker-compose up --build --no-deps -d --force-recreate
	# Fix race condition here between container starting and test running.
	go test ./test/... || true
	cd test &&  docker-compose down