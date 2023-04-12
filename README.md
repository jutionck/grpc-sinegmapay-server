# SinegmaPay gRPC Server

## Generate `.Proto`
```bash
protoc --go_out=./service --go-grpc_out=./service model/sinegmapay.proto
```

## Run
```bash
GRPC_URL=localhost:1200 go run . 
```