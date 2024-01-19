#Migrations
```
migrate create -ext sql -dir postgres/migrations [table]
source .env ====> (export POSTGRESQL_URL=postgres://postgres:admin@localhost:5432/share_room?sslmode=disable)
migrate -path "postgres/migrations" -database "$POSTGRESQL_URL" up  
```
