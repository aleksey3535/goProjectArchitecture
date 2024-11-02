package repository

import "github.com/jmoiron/sqlx"

type PartOfTheAppOne interface {
	DoSomething() error
}

type PartOfTheAppTwo interface {
	DoSomething() error
}

type Repository struct {
	PartOfTheAppOne
	PartOfTheAppTwo
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		PartOfTheAppOne: NewPartOne(db),
		PartOfTheAppTwo: NewPartTwo(db),
	}

}