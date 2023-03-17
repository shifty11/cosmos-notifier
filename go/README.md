# Go backend

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

# generate migrations
go run main.go create-migrations

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
```
