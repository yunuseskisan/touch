.PHONY: test

SERVICE = service.rpc.assets
CI_COMMIT_SHA ?= latest

generate:
	cd ../ && bin/generate_go_proto $(SERVICE)

build:
	cd ../ && bin/build_service $(SERVICE) $(CI_COMMIT_SHA)
