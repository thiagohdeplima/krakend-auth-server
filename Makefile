test: mock
	go test -v ./...

mock:
	@go install github.com/vektra/mockery/v2@v2.36.0
	@~/go/bin/mockery --all

build:
	docker-compose up builder

run: build
	docker-compose up -d

reload: build
	docker-compose restart server
