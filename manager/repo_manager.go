package manager

import "github.com/jutionck/grpc-sinegmapay-server/repository"

type RepositoryManager interface {
	SinegmaPayRepo() repository.SinegmaPayRepository
}

type repositoryManager struct {
	sinegmaPayRepo repository.SinegmaPayRepository
}

func (r *repositoryManager) SinegmaPayRepo() repository.SinegmaPayRepository {
	return r.sinegmaPayRepo
}

func NewRepositoryManager() RepositoryManager {
	repo := new(repositoryManager)
	repo.sinegmaPayRepo = repository.NewSinegmaPayRepository()
	return repo
}
