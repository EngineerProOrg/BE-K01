DB_URL=mysql://root:secret@tcp(127.0.0.1:3306)/simple_school

.PHONY: mysql

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

migrate:
	migrate create -ext sql -dir db/migration $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" --verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" --verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" --verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" --verbose down 1
