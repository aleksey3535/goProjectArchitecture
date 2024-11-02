package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type PartOne struct {
	db *sqlx.DB
}

func NewPartOne(db *sqlx.DB) *PartOne {
	return &PartOne{db: db}
}
func (p1 *PartOne) DoSomething() error {
	log.Print("Some work with database")
	return nil
}