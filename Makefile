mysql/init:
	docker-compose -f db/database.yaml up --detach
	docker rm imageboard_db_waiter

mysql/down:
	docker-compose -f db/database.yaml down

compose/up:
	@go mod vendor
	docker-compose up -d

compose/down:
	docker-compose down

up: mysql/init compose/up

down: compose/down mysql/down
