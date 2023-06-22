rm -rf ./dist
go build -o ./dist/server ./cmd/server/main.go 
go build -o ./dist/client ./cmd/client/main.go 
