# note: call scripts from /scripts

# make new_migration MESSAGE_NAME="message name"
new_migration:
	migrate create -ext sql -dir scripts/migration/ -seq $(MESSAGE_NAME)
up_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local" -verbose up
down_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local" -verbose down
proto_aap:
	protoc --proto_path=pkg/types/proto/ --go_out=pkg/types/proto/pb/authen_and_post --go_opt=paths=source_relative \
        --go-grpc_out=pkg/types/proto/pb/authen_and_post --go-grpc_opt=paths=source_relative \
        authen_and_post.proto
proto_newsfeed:
	protoc --proto_path=pkg/types/proto/ --go_out=pkg/types/proto/pb/newsfeed --go_opt=paths=source_relative \
        --go-grpc_out=pkg/types/proto/pb/newsfeed --go-grpc_opt=paths=source_relative \
        newsfeed.proto
tidy:
	go mod tidy
.PHONY: vendor
vendor:
	go mod vendor -v
docker_clear:
	docker volume rm $(docker volume ls -qf dangling=true) & docker rmi $(docker images -f "dangling=true" -q)
compose_up_rebuild:
	docker compose up --build --force-recreate
compose_up:
	docker compose up
