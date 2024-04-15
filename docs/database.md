# DB

Access the local development DB:
```bash
sqlite3 db/sqlite3.db
```

## Migrations

See [golang-migrate](https://github.com/golang-migrate/migrate/tree/master) for additional documentation.

[Command installation](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```sh
export DB_URL="sqlite3://db/sqlite3.db"
```

To create a migration:
```bash
migrate create -ext sql -dir db/migrations -seq sample_migration_name
```

To run up migrations:
```bash
migrate -database ${DB_URL} -path db/migrations up
```

To drop the db:
```bash
migrate -database ${DB_URL} -path db/migrations drop
```

After a migration error is encountered, the DB is marked dirty and a migration version must be forced before any more migrations can be run:
```bash
migrate -database ${DB_URL} -path db/migrations force <DB VERSION BEFORE ERROR>
```