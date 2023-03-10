build:
	docker run -it -w /app \
		-v "${PWD}:/app" \
		krakend/builder:2.2.1 \
			go build -buildmode=plugin -o authorization-server.so .

run: build
	docker-compose up -d

reload: build
	docker-compose restart krakend
	
