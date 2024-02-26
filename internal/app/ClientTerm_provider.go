package app

import (
	"log"

	clientterm "github.com/Alphonnse/noteTakingApp/internal/api/clientTerm"
	"github.com/Alphonnse/noteTakingApp/internal/config"
	"github.com/Alphonnse/noteTakingApp/internal/repository"
	clientTermRepository "github.com/Alphonnse/noteTakingApp/internal/repository/clientTerm"
	"github.com/Alphonnse/noteTakingApp/internal/service"
	clientTermService "github.com/Alphonnse/noteTakingApp/internal/service/clientTerm"
)

type clientTermProvider struct {
	grpcConfig        config.GRPCConfig
	clientTermRepo    repository.ClientTermRepository
	clientTermService service.ClientTermService
	clientTermImpl    *clientterm.Implemetation
}

func NewClientTermProvider() *clientTermProvider {
	return &clientTermProvider{}
}

func (provider *clientTermProvider) GRPCConfig() config.GRPCConfig {
	if provider.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("Error when getting grpc config: %s", err.Error())
		}
		provider.grpcConfig = cfg
	}

	return provider.grpcConfig
}

func (provider *clientTermProvider) ClientTermRepo() repository.ClientTermRepository {
	if provider.clientTermRepo == nil {
		provider.clientTermRepo = clientTermRepository.NewRepository()
	}

	return provider.clientTermRepo
}

func (provider *clientTermProvider) ClientTermService() service.ClientTermService {
	if provider.clientTermService == nil {
		provider.clientTermService = clientTermService.NewService(
			provider.ClientTermRepo(),
		)
	}
	
	return provider.clientTermRepo
}

func (provider *clientTermProvider) ClientTermImpl() *clientterm.Implemetation {
	if provider.clientTermImpl == nil {
		provider.clientTermImpl = clientterm.NewImplementation(
			provider.ClientTermService(),
		)
	}

	return provider.clientTermImpl
}
