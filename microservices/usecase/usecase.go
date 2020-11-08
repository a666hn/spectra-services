package usecase

import (
	"github.com/skinnyguy/spectra-services/microservices/repository"
)

type UseCase struct {
	repository *repository.AbstractRepository
}

func NewUseCase(repo *repository.AbstractRepository) *UseCase {
	return &UseCase{
		repository: repo,
	}
}
