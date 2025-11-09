# Go Template

Template for a Go service

Setup:
```
make dep
```

Run with docker:
```
docker-compose up -d
```

Run without docker:
```
docker-compose up -d postgres # run only postgres via docker
make run
```

Lint:
```
make lint
```

Generate db queries:
```
make sqlc
```

Create a db migration:
```
make migrate/create name=<migration_name>
```
