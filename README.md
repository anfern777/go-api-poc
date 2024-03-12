# web3-gaming

## Migration UP

```sh
$ migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
```

## Migration DOWN

```sh
$ migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down
```

## Resolve migration errors
It is necessary to determine if the migration was applied partially or not at all. Once you've determined this, the database version should be forced to reflect its true state using the force command.
```sh
$ migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <VERSION>
```