APP_NAME?=web

server:
	go run cmd/web/main.go

build:
	$(MAKE) -C cmd/$(APP_NAME) build

d.up:
	docker-compose up 

d.down:
	docker-compose down

d.up.build:
	docker-compose --build up