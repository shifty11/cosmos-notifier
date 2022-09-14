# DaoDao Notifier
DaoDao Notifier is a subscription service that notifies about new governance proposals of [DaoDao's](https://daodao.zone/).

STATUS: ALPHA VERSION - NOT READY FOR PRODUCTION

## Setup

Full development setup
```bash
cp .env.template .env   # make changes to env
docker-compose -f dev.yml up
```
You need to create a Telegram and a Discord bot and get the token. 
You also need to get the Discord OAuth2 client id and secret.
They are all needed in the .env file.

## Database

### Docker
run psql-db and pgadmin4
```bash
docker-compose -f local-db.yml up   # add -d for detached mode
```
remove containers and volumes
```bash
docker-compose -f local-db.yml down -v
```

### Install cli

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest migrate -database "postgres:
```

### Generate models

Edit ent/schema/*.go and run:

```bash
go generate ./ent
```

### Create migrations

```bash
go run main.go create-migrations
```
combined (generate models and migrations)
    
```bash
go generate ./ent && go run main.go create-migrations
```

### Apply migrations

```bash
migrate -source file://database/migrations -database "postgres://postgres:postgres@localhost:5432/daodao-notifier-db?sslmode=disable&TimeZone=Europe/Zurich" up
```


## API

### Protoc Installation
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Frontend

```bash
envoy -c envoy.yaml
```