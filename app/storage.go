package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/domain/user"
	"github.com/sealoftime/adapteris/pgx"
)

type Storages struct {
	connPool *pgxpool.Pool
	accounts user.Repository
	schools  school.SchoolRepository
	stages   school.StageRepository
	steps    school.StepRepository
	events   school.EventRepository
}

func (a *App) initStorage() {
	a.Storage.connPool = connectTo(a.Config.DbUri)
	a.Storage.accounts = pgx.NewAccountStorage(a.Log)
	a.Storage.schools = pgx.NewSchoolStorage(a.Log)
	a.Storage.stages = pgx.NewStageStorage(a.Log)
	a.Storage.steps = pgx.NewStepStorage(a.Log)
	a.Storage.events = pgx.NewEventStorage(a.Log)
}

// Open new connection
func connectTo(databaseUrl string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("couldn't connec	t to databse: %+v\n", err)
	}

	return pool
}
