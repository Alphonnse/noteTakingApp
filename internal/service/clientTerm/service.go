package clientterm

import "github.com/Alphonnse/noteTakingApp/internal/repository"

type service struct {
	clientTermRepository repository.ClientTermRepository
}

func NewService(clientRepository repository.ClientTermRepository) *service {
	return &service{
		clientTermRepository: clientRepository,
	}
}
