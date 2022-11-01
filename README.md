# DaoDao Notifier
DaoDao Notifier is a subscription service that notifies about new governance proposals of [DaoDao's](https://daodao.zone/).

## Setup

Full development setup
```bash
cp .env.template .env   # make changes to env
docker-compose -f dev.yml up
```
You need to create a Telegram and a Discord bot and get the token. 
You also need to get the Discord OAuth2 client id and secret.
They are all needed in the .env file.

If you just want to run the whole application you can use the dev docker-compose file. 
Otherwise follow the steps below.
```bash
docker-compose -f dev.yml up
```

## Database

### Docker
```bash
# run psql-db and pgadmin4
docker-compose -f local-db.yml up   # add -d for detached mode

# remove containers and volumes
docker-compose -f local-db.yml down -v
```

### Install cli

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest migrate -database "postgres:
```

### Generate models and migrations

Edit ent/schema/*.go and run:

```bash
# generate models
go generate ./ent

# generate migrations
go run main.go create-migrations

# or combined
go generate ./ent && go run main.go create-migrations
```

### Apply migrations

```bash
# install cli
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest migrate -database "postgres:

# apply migrations
migrate -source file://database/migrations -database "postgres://postgres:postgres@localhost:5432/cosmos-notifier-db?sslmode=disable&TimeZone=Europe/Zurich" up
```


## API

### Protoc Installation
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Frontend

**[Flutter](webapp/README.md)**