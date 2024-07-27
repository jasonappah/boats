app-dev:
  wails dev -s

pb *ARGS:
  go run cmd/pb/main.go {{ARGS}}

db:
  open ./pb_data/data.db