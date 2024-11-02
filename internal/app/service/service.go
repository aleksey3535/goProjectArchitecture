package service

import "app/internal/app/repository"

type PartOfTheAppOneService interface {
	DoSomething() error
}

type PartOfTheAppTwoService interface {
	DoSomething() error
}

type Service struct {
	PartOfTheAppOneService
	PartOfTheAppTwoService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		PartOfTheAppOneService: NewPartOneService(repo),
		PartOfTheAppTwoService: NewPartTwoService(repo),

	}
}