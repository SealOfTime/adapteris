package pgx

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

const (
	AverageEventsInStep = 2
)

var (
	Step = struct {
		Id, MustComplete, StageId string
	}{
		Id:           "id",
		MustComplete: "must_complete",
		StageId:      "stage_id",
	}
	allStepColumns = columns{
		Step.Id,
		Step.MustComplete,
		Step.StageId,
	}
	selectStepById = fmt.Sprintf(
		"SELECT %s FROM step WHERE %s=$1",
		allStepColumns.sqlString(), Step.Id,
	)
	selectEventsForStep = fmt.Sprintf(
		"SELECT %s FROM event WHERE %s = $1",
		allEventColumns.sqlString(), Event.StepId,
	)
)

type StepStorage struct {
	log log.Logger
}

func NewStepStorage(log log.Logger) *StepStorage {
	return &StepStorage{
		log: log,
	}
}

var _ school.StepRepository = (*StepStorage)(nil)

func (s *StepStorage) FindById(ctx context.Context, id int64) (*school.Step, error) {
	tx := TxFromCtx(ctx)

	step, err := mapStep(tx.QueryRow(ctx, selectStepById, id))
	if err != nil {
		return nil, err
	}

	return &step, nil
}

func (s *StepStorage) Save(ctx context.Context, saved *school.Step) (returned *school.Step, err error) {
	tx := TxFromCtx(ctx)

	query, vals := buildUpsertStepQuery(*saved, 0)
	s.log.Log(query)
	dbStep, err := mapStep(tx.QueryRow(ctx, query, vals...))
	if err != nil {
		return nil, err
	}

	b := &pgx.Batch{}
	for _, event := range saved.Events {
		query, params := buildUpsertEventQuery(event, dbStep.Id)
		b.Queue(query, params...)
	}
	br := tx.SendBatch(ctx, b)
	defer func() {
		if brErr := br.Close(); brErr != nil {
			s.log.Log("error closing batch response: %+v", brErr)
			if err == nil {
				saved, err = nil, brErr
			}
		}
	}()

	dbStep.Events = make([]school.Event, len(saved.Events))
	for i := range dbStep.Events {
		if dbStep.Events[i], err = mapEvent(br.QueryRow()); err != nil {
			return nil, err
		}
	}

	return &dbStep, nil
}

func buildUpsertStepQuery(step school.Step, stageId int64) (string, []sqlValue) {
	if step.Id == 0 {
		return buildInsertStepQuery(step, stageId)
	}
	return buildUpdateStepQuery(step, stageId)
}

func buildInsertStepQuery(step school.Step, stageId int64) (string, []sqlValue) {
	if stageId == 0 {
		panic("can't create step without a link to a stage")
	}

	insert := sqlValues{
		Step.MustComplete: step.MustComplete,
		Step.StageId:      stageId,
	}

	cols, vals := insert.split()
	return fmt.Sprintf(
		`INSERT INTO step (%s) VALUES (%s) RETURNING %s`,
		cols.sqlString(), cols.sqlParams(), allStepColumns.sqlString(),
	), vals
}

func buildUpdateStepQuery(step school.Step, stageId int64) (string, []sqlValue) {
	update := sqlValues{
		Step.MustComplete: step.MustComplete,
	}

	if stageId != 0 {
		update[Step.StageId] = stageId
	}

	cols, vals := update.split()
	updatedCols := make([]colName, 0, len(cols))
	for i, col := range cols {
		updatedCols = append(updatedCols, fmt.Sprintf("%s=$%d", col, i+1))
	}
	return fmt.Sprintf(
		`UPDATE step SET %s WHERE %s=%d RETURNING %s`,
		strings.Join(updatedCols, ", "), Step.Id, step.Id, allStepColumns.sqlString(),
	), vals
}

func mapSteps(rows pgx.Rows) ([]school.Step, error) {
	res := make([]school.Step, 0, AverageStepsInStage)
	for rows.Next() {
		step, err := mapStep(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, step)
	}
	return res, nil
}
func mapStep(row pgx.Row) (s school.Step, err error) {
	var unused interface{}
	err = row.Scan(
		&s.Id,
		&s.MustComplete,
		&unused,
	)
	return
}
