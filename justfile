set dotenv-load

mock:
    mockgen -source=go/database/telegram_chat_manager.go -destination=go/database/mock_types/telegram_chat_manager.go
    mockgen -source=go/database/discord_channel_manager.go -destination=go/database/mock_types/discord_channel_manager.go

generate-protobufs:
    protoc -I=api/ --go_out=go/services/grpc/protobuf/ --go_opt=module=github.com/shifty11/cosmos-notifier/services/grpc/protobuf \
            --go-grpc_out=go/services/grpc/protobuf/ --go-grpc_opt=module=github.com/shifty11/cosmos-notifier/services/grpc/protobuf \
            --dart_out=grpc:webapp/lib/api/protobuf/dart/ api/*.proto && \
    protoc -I=api/ --dart_out=grpc:webapp/lib/api/protobuf/dart/ google/protobuf/timestamp.proto google/protobuf/empty.proto google/protobuf/duration.proto

generate-models:
    cd go && go generate ./ent

visualize-models:
    cd go && go run -mod=mod ariga.io/entviz ./ent/schema/

generate-migrations:
    cd go && go run main.go database create-migrations

migrate:
    cd go && go run main.go database migrate

test-go:
    cd go && go test ./...

install-web:
    cd web
    cargo install --locked trunk
    nvm use 18 && npm install -D tailwindcss

format-web:
    cd web && cargo fmt && cargo clippy

web:
    cd web && trunk serve
