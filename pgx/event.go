package pgx

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

var (
	Event = struct {
		Id, Name, Description, StepId string
	}{
		Id:          "id",
		Name:        "name",
		Description: "description",
		StepId:      "step_id",
	}
	allEventColumns = columns{
		Event.Id,
		Event.Name,
		Event.Description,
		Event.StepId,
	}
)

type EventStorage struct {
	log log.Logger
}

var _ school.EventRepository = (*EventStorage)(nil)

func NewEventStorage(
	log log.Logger,
) *EventStorage {
	return &EventStorage{
		log: log,
	}
}

func (s *EventStorage) FindById(ctx context.Context, id int64) (*school.Event, error) {
	panic("todo")
}

func (s *EventStorage) Save(ctx context.Context, saved *school.Event) (returned *school.Event, err error) {
	panic("todo")
}

func buildUpsertEventQuery(event school.Event, stepId int64) (string, []sqlValue) {
	if event.Id == 0 && stepId == 0 {
		panic("can't create new event without link to a step")
	}

	insert := sqlValues{
		Event.Name: event.Name,
	}

	if event.Description != nil {
		insert[Event.Description] = event.Description
	}

	if stepId != 0 {
		insert[Event.StepId] = stepId
	}

	cols, vals := insert.split()
	updatedCols := make([]string, len(cols))
	for i, col := range cols {
		updatedCols[i] = fmt.Sprintf("%[1]s=EXCLUDED.%[1]s", col)
	}
	return fmt.Sprintf(
		`INSERT INTO event (%s) VALUES (%s) 
		ON CONFLICT (id) DO UPDATE 
			SET %s
		RETURNING %s`,
		cols.sqlString(), cols.sqlParams(), strings.Join(updatedCols, ", "), allEventColumns.sqlString(),
	), vals
}

func mapEvents(rows pgx.Rows) ([]school.Event, error) {
	events := make([]school.Event, 0, AverageEventsInStep)

	for rows.Next() {
		event, err := mapEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func mapEvent(row pgx.Row) (event school.Event, err error) {
	var ignored interface{}
	if err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&ignored,
	); err != nil {
		return
	}
	return
}
