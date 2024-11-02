package service

import "app/internal/app/repository"

type PartTwoService struct {
	repo *repository.Repository
}

func NewPartTwoService(repo *repository.Repository) *PartTwoService {
	return &PartTwoService{repo:repo}
}

func(s2 *PartTwoService) DoSomething() error {
	return s2.repo.PartOfTheAppTwo.DoSomething()
}