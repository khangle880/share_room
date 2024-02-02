#Migrations
```
migrate create -ext sql -dir postgres/migrations [table]
source .env ====> (export POSTGRESQL_URL=postgres://postgres:admin@localhost:5432/share_room?sslmode=disable)
migrate -path "postgres/migrations" -database "$POSTGRESQL_URL" up  
```

```DataLoader
go run github.com/vektah/dataloaden UserLoader string *github.com/dataloaden/example.User
go run github.com/vektah/dataloaden UserLoader 'github.com/google/uuid.UUID' '*github.com/khangle880/share_room/pg.User'
go run github.com/vektah/dataloaden ProfileLoader 'github.com/google/uuid.UUID' '*github.com/khangle880/share_room/pg.Profile'
```