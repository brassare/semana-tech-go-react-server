package gen

//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yaml
//go:generate go run ./cmd/tools/terndotenv/main.go
