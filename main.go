package main

import "github.com/jutionck/grpc-sinegmapay-server/delivery"

func main() {
	delivery.Server().Run()
}
