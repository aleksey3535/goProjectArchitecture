package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type PartTwo struct {
	db *sqlx.DB
}

func NewPartTwo(db *sqlx.DB) *PartTwo {
	return &PartTwo{db: db}
}

func(p2 *PartTwo) DoSomething() error {
	log.Print("do something with database v2")
	return nil
}