# Go backend

## Services

```bash
source ../.env.dev  # set environment variables
```

| Service           | Command                                  | Description                        |
|-------------------|:-----------------------------------------|------------------------------------|
| Telegram          | go run main.go service telegram          | Telegram bot                       |
| Discord           | go run main.go service discord           | Discord bot                        |
| GRPC              | go run main.go service grpc              | GRPC server                        |
| Chain Crawler     | go run main.go service chain-crawler     | Crawler to update chains           |
| Contract Crawler  | go run main.go service contract-crawler  | Crawler to update DAODAO contracts |
| Validator Crawler | go run main.go service validator-crawler | Crawler to update validators       |

## Database

### Install CLI

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Generate models and migrations

Edit ent/schema/*.go and run:

```bash
# generate models
go generate ./ent
```
```bash
# generate migrations
go run main.go create-migrations
```
```bash
# or combined
go generate ./ent && go run main.go create-migrations
```

### Apply migrations

```bash
migrate -source file://go/database/migrations -database "postgres://postgres:postgres@localhost:5432/cosmos-notifier-db?sslmode=disable&TimeZone=Europe/Zurich" up
```

### Create schema visualization

```bash
go run -mod=mod ariga.io/entviz ./ent/schema/
```

## API

### Protoc Installation

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
