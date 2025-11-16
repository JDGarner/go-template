# Go Template

### Template for a Go service with:
- http server (using github.com/labstack/echo/v4)
- postgres (using github.com/jackc/pgx/v5)
- sqlc (sql query/model generator)
- golang-migrate
- testcontainers-go

#### When forking this repo, replace 'go-template' references the following files:
- `go.mod`
- `Makefile`
- `Dockerfile`
- `docker-compose.yml`
- `golangci.yml`
- `internal/tests/docker-compose.yml`
- `internal/store/migrations/*.sql`

And:
- Update `store/schema.sql` with a new schema
- Create .env file based on .env.example

#### Setup:
```
make dep
```

#### Run with docker:
```
docker-compose up -d
```

#### Run without docker:
```
docker-compose up -d postgres # run only postgres via docker
make run
```

#### Lint:
```
make lint
```

#### Tests:
- Ensure the following env vars are set if using docker with colima:
```
set:
export DOCKER_HOST=unix://${HOME}/.colima/default/docker.sock
export TESTCONTAINERS_RYUK_DISABLED=true
```
Then run:
```
make test
```

#### Generate db queries:
```
make sqlc
```

#### Create a db migration:
```
make migrate/create name=<migration_name>
```
