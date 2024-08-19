proto:
	protoc --go_out=. pkg/**/pb/*.proto
	protoc --go-grpc_out=. pkg/**/pb/*.proto

server:
	go run cmd/main.go