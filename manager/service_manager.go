package manager

import "github.com/jutionck/grpc-sinegmapay-server/service"

type ServiceManager interface {
	SinegmaPayService() *service.SinegmaPayService
}

type serviceManager struct {
	sinegmaPayService *service.SinegmaPayService
}

func (s *serviceManager) SinegmaPayService() *service.SinegmaPayService {
	return s.sinegmaPayService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	srv := new(serviceManager)
	srv.sinegmaPayService = service.NewSinegmaPayService(repoManager.SinegmaPayRepo())
	return srv
}
