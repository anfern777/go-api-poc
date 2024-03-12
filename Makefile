migration_up: migrate -path database/migration/ -database "postgresql://postgres:test123@localhost:5432/postgres?sslmode=disable" -verbose up

migration_down: migrate -path database/migration/ -database "postgresql://postgres:test123@localhost:5432/postgres?sslmode=disable" -verbose down

migration_fix: migrate -path database/migration/ -database "postgresql://postgres:test123@localhost:5432/postgres?sslmode=disable" force 1