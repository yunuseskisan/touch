.PHONY: test

test:
	docker-compose up --build --no-deps -d --force-recreate
	# Fix race condition here between container starting and test running.
	go test ./test/... || true
	docker-compose down