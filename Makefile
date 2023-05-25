# note: call scripts from /scripts

# make new_migration MESSAGE_NAME=test
new_migration:
	migrate create -ext sql -dir scripts/migration/ -seq $(MESSAGE_NAME)
up_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local" -verbose up
down_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local" -verbose down