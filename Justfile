app-dev:
  wails dev -s


# https://pocketbase.io/docs/go-migrations/
pb *ARGS:
  go run cmd/pb/main.go {{ARGS}}

db:
  open ./pb_data/data.db