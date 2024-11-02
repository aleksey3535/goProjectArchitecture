package service

import "app/internal/app/repository"

type PartOneService struct {
	repo *repository.Repository
}

func NewPartOneService(repo *repository.Repository) *PartOneService {
	return &PartOneService{repo}
}

func(s1 PartOneService) DoSomething() error {
	return s1.repo.PartOfTheAppOne.DoSomething()
}