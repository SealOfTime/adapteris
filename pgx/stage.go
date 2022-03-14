package pgx

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

type StageStorage struct {
	log log.Logger
}

func NewStageStorage(log log.Logger) *StageStorage {
	return &StageStorage{
		log: log,
	}
}

var _ school.StageRepository = (*StageStorage)(nil)

var (
	Stage = struct {
		Id, SchoolId, Name, Description, StartDate, EndDate string
	}{
		Id:          "id",
		Name:        "name",
		Description: "description",
		StartDate:   "start_date",
		EndDate:     "end_date",
		SchoolId:    "school_id",
	}
	allStageColumns = columns{
		Stage.Id,
		Stage.Name,
		Stage.Description,
		Stage.StartDate,
		Stage.EndDate,
		Stage.SchoolId,
	}
)

var (
	selectStageById = fmt.Sprintf(
		"SELECT %s FROM stage WHERE %s=$1",
		allStageColumns.sqlString(), Stage.Id,
	)
	selectStepsForStage = fmt.Sprintf(
		"SELECT %s FROM step WHERE %s=$1",
		allStepColumns.sqlString(), Step.StageId,
	)
)

const (
	AverageStepsInStage = 3
)

func (s *StageStorage) FindById(ctx context.Context, id int64) (*school.Stage, error) {
	tx := TxFromCtx(ctx)

	dbStage, err := mapStage(tx.QueryRow(ctx, selectStageById, id))
	if err != nil {
		return nil, fmt.Errorf("error finding stage by id: %w", err)
	}

	if err := s.findRelationships(ctx, &dbStage); err != nil {
		return nil, fmt.Errorf("error finding relationships for stage: %w", err)
	}

	return &dbStage, nil
}

func (s *StageStorage) findRelationships(ctx context.Context, stage *school.Stage) (err error) {
	tx := TxFromCtx(ctx)

	rows, err := tx.Query(ctx, selectStepsForStage, stage.Id)
	if err != nil {
		return fmt.Errorf("error finding steps for stage: %w", err)
	}

	stage.Steps, err = mapSteps(rows)
	if err != nil {
		return fmt.Errorf("error mapping steps for stage: %w", err)
	}

	b := &pgx.Batch{}
	for _, step := range stage.Steps {
		b.Queue(selectEventsForStep, step.Id)
	}
	if err = processBatch(tx.SendBatch(ctx, b), func(br pgx.BatchResults) error {
		var (
			rows pgx.Rows
			err  error
		)
		for _, step := range stage.Steps {
			if rows, err = br.Query(); err != nil {
				return fmt.Errorf("error quering events for step(id=%d): %w", step.Id, err)
			}
			if step.Events, err = mapEvents(rows); err != nil {
				return fmt.Errorf("error mapping events for steps: %w", err)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *StageStorage) Save(ctx context.Context, saved *school.Stage) (returned *school.Stage, err error) {
	tx := TxFromCtx(ctx)

	query, params := buildUpsertStageSql(*saved, 0)
	s.log.Log(query)
	dbStage, err := mapStage(tx.QueryRow(ctx, query, params...))
	if err != nil {
		return nil, err
	}

	b := &pgx.Batch{}
	for _, step := range saved.Steps {
		query, params := buildUpsertStepQuery(step, dbStage.Id)
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

	dbStage.Steps = make([]school.Step, len(saved.Steps))
	for i := range dbStage.Steps {
		if dbStage.Steps[i], err = mapStep(br.QueryRow()); err != nil {
			return nil, err
		}
	}

	return &dbStage, nil
}

func buildUpsertStageSql(stage school.Stage, schoolId int64) (string, []sqlValue) {
	if stage.Id == 0 {
		return buildInsertStageQuery(stage, schoolId)
	}

	return buildUpdateStageQuery(stage, schoolId)
}

func buildInsertStageQuery(stage school.Stage, schoolId int64) (string, []sqlValue) {
	if schoolId == 0 {
		panic("can't create stage without link to a school")
	}
	insert := sqlValues{
		Stage.Name:     stage.Name,
		Stage.SchoolId: schoolId,
	}

	if stage.StartDate != nil {
		insert[Stage.StartDate] = stage.StartDate
	}

	if stage.EndDate != nil {
		insert[Stage.EndDate] = stage.EndDate
	}

	cols, vals := insert.split()
	return fmt.Sprintf(
		`INSERT INTO stage (%s) VALUES (%s) RETURNING %s`,
		cols.sqlString(), cols.sqlParams(), allStageColumns.sqlString(),
	), vals
}

func buildUpdateStageQuery(stage school.Stage, schoolId int64) (string, []sqlValue) {
	update := sqlValues{
		Stage.Name: stage.Name,
	}

	if stage.StartDate != nil {
		update[Stage.StartDate] = stage.StartDate
	}

	if stage.EndDate != nil {
		update[Stage.EndDate] = stage.EndDate
	}

	if schoolId != 0 {
		update[Stage.SchoolId] = schoolId
	}

	cols, vals := update.split()
	updatedCols := make([]colName, 0, len(cols))
	for i, col := range cols {
		updatedCols = append(updatedCols, fmt.Sprintf("%s=$%d", col, i+1))
	}

	return fmt.Sprintf(
		`UPDATE stage SET %s WHERE %s=%d RETURNING %s`,
		strings.Join(updatedCols, ", "), Stage.Id, stage.Id, allStageColumns.sqlString(),
	), vals
}

func mapStages(rows pgx.Rows) ([]school.Stage, error) {
	res := make([]school.Stage, 0, AverageStagesInSchool)

	for rows.Next() {
		stage, err := mapStage(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, stage)
	}

	return res, nil
}
func mapStage(row pgx.Row) (s school.Stage, err error) {
	var description *string
	var schoolIdDump int64
	err = row.Scan(
		&s.Id,
		&s.Name,
		&description,
		&s.StartDate,
		&s.EndDate,
		&schoolIdDump,
	)

	if description != nil {
		s.Description = *description
	}

	return
}
