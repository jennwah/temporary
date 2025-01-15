package usecase

import (
	"github.com/jennwah/ryde-backend-engineer/internal/usecase/getnearbyfriends/internal/contract"
)

type UseCase struct {
	repo contract.Repository
}

func New(repo contract.Repository) UseCase {
	return UseCase{
		repo: repo,
	}
}
