package handler

import (
	"sync"

	"github.com/skinnyguy/spectra-services/microservices/repository"
	"github.com/skinnyguy/spectra-services/microservices/usecase"
)

var uc *usecase.UseCase
var oneUc sync.Once

func GetUsecase() *usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(getRepository())
	})

	return uc
}

var repo *repository.AbstractRepository
var repoOnce sync.Once

func getRepository() *repository.AbstractRepository {
	repoOnce.Do(func() {
		repo = repository.NewRepository()
	})

	return repo
}
