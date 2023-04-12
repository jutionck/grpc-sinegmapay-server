package delivery

import (
	"github.com/jutionck/grpc-sinegmapay-server/config"
	"github.com/jutionck/grpc-sinegmapay-server/manager"
	"github.com/jutionck/grpc-sinegmapay-server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SinegmaPayServer struct {
	listen         net.Listener
	server         *grpc.Server
	serviceManager manager.ServiceManager
}

func (s *SinegmaPayServer) Run() {
	service.RegisterSinegmaPaymentServer(s.server, s.serviceManager.SinegmaPayService())
	log.Println("Server run:", s.listen.Addr().String())
	err := s.server.Serve(s.listen)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func Server() *SinegmaPayServer {
	sinegmaPayServer := new(SinegmaPayServer)
	c := config.NewConfig()
	list, err := net.Listen("tcp", c.Url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	repoManager := manager.NewRepositoryManager()
	sinegmaPayServer.listen = list
	sinegmaPayServer.server = grpcServer
	sinegmaPayServer.serviceManager = manager.NewServiceManager(repoManager)
	return sinegmaPayServer
}
