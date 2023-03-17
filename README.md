# Cosmos Notifier
Cosmos Notifier is a subscription service that notifies about new governance proposals of [Cosmos chains](https://cosmos.network/) and [DaoDao's](https://daodao.zone/).

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
## Backend

**[Go](go/README.md)**

## Frontend

**[Flutter](webapp/README.md)**
