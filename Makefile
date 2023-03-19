test: mock
	go test -v ./...

mock:
ifeq ($(shell test -e ~/go/bin/mockery1 && echo -n "yes"), "yes")
go install github.com/vektra/mockery/v2
endif

	@~/go/bin/mockery --all

build:
	docker-compose up builder

run: build
	docker-compose up -d

reload: build
	docker-compose restart server
