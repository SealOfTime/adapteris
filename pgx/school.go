package pgx

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

var (
	School = struct {
		id, name, visible, start, end, registrationStart, registrationEnd string
	}{
		id:                "id",
		name:              "name",
		visible:           "visible",
		start:             "start_date",
		end:               "end_date",
		registrationStart: "register_start",
		registrationEnd:   "register_end",
	}
	allSchoolColumns = columns{
		School.id,
		School.name,
		School.visible,
		School.start,
		School.end,
		School.registrationStart,
		School.registrationEnd,
	}
	selectActiveSchool = fmt.Sprintf(
		"SELECT %s FROM school WHERE %s=true",
		allSchoolColumns.sqlString(), School.visible,
	)
	selectSchoolByIdSql = fmt.Sprintf(
		"SELECT %s FROM school WHERE %s=$1",
		allSchoolColumns.sqlString(), School.id,
	)
	selectStagesForSchoolQuery = fmt.Sprintf(
		"SELECT %s FROM stage WHERE %s=$1",
		allStageColumns.sqlString(), Stage.SchoolId,
	)
)

type SchoolStorage struct {
	log log.Logger
}

func NewSchoolStorage(log log.Logger) *SchoolStorage {
	return &SchoolStorage{
		log: log,
	}
}

var _ school.SchoolRepository = (*SchoolStorage)(nil)

func (s *SchoolStorage) FindById(ctx context.Context, id int64) (*school.School, error) {
	tx := TxFromCtx(ctx)

	sch, err := mapSchool(tx.QueryRow(
		ctx,
		selectSchoolByIdSql,
		id,
	))
	if err != nil {
		return nil, err
	}

	if err = s.findRelationships(ctx, &sch); err != nil {
		return nil, fmt.Errorf("error finding relationships for school: %w", err)
	}
	return &sch, nil
}

const (
	//AverageStagesInSchool is a constant of average stages per year
	AverageStagesInSchool int = 4
)

func (s *SchoolStorage) findRelationships(ctx context.Context, sch *school.School) error {
	tx := TxFromCtx(ctx)

	rows, err := tx.Query(ctx, selectStagesForSchoolQuery, sch.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			//Absolutely normal scenario to have no stages for a school
			return nil
		}
		return fmt.Errorf("error querying stages: %w", err)
	}

	if sch.Stages, err = mapStages(rows); err != nil {
		return fmt.Errorf("error mapping stages: %w", err)
	}

	b := &pgx.Batch{}
	for _, stage := range sch.Stages {
		b.Queue(selectStepsForStage, stage.Id)
	}
	if err = processBatch(tx.SendBatch(ctx, b), func(br pgx.BatchResults) error {
		for i := range sch.Stages {
			rows, err = br.Query()
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					//It's okay to not have steps in a stage
					continue
				}
				return fmt.Errorf("error querying steps: %w", err)
			}
			if sch.Stages[i].Steps, err = mapSteps(rows); err != nil {
				return fmt.Errorf("error mapping steps: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error processing steps batch: %w", err)
	}

	b = &pgx.Batch{}
	for _, stage := range sch.Stages {
		for _, step := range stage.Steps {
			b.Queue(selectEventsForStep, step.Id)
		}
	}
	if err = processBatch(tx.SendBatch(ctx, b), func(br pgx.BatchResults) error {
		var (
			rows pgx.Rows
			err  error
		)
		for stgIdx := range sch.Stages {
			for stpIdx := range sch.Stages[stgIdx].Steps {
				if rows, err = br.Query(); err != nil {
					if errors.Is(err, pgx.ErrNoRows) {
						continue //it's okay to not have events in a step
					}
					return fmt.Errorf("error querying events for step: %w", err)
				}
				if sch.Stages[stgIdx].Steps[stpIdx].Events, err = mapEvents(rows); err != nil {
					return fmt.Errorf("error mapping events for step: %w", err)
				}
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error processing events batch: %w", err)
	}

	return nil
}

func (s *SchoolStorage) Save(ctx context.Context, sch *school.School) (saved *school.School, err error) {
	tx := TxFromCtx(ctx)

	query, params := buildUpsertSchoolSql(sch)
	s.log.Log(query)

	dbSchool, err := mapSchool(tx.QueryRow(ctx, query, params...))
	if err != nil {
		return nil, err
	}

	b := &pgx.Batch{}
	for _, stage := range sch.Stages {
		query, params := buildUpsertStageSql(stage, dbSchool.Id)
		s.log.Log(query)
		b.Queue(query, params...)
	}
	br := tx.SendBatch(ctx, b)
	defer func() {
		if brErr := br.Close(); brErr != nil {
			saved, err = nil, brErr
		}
	}()

	dbSchool.Stages = make([]school.Stage, len(sch.Stages))
	for i := range dbSchool.Stages {
		dbSchool.Stages[i], err = mapStage(br.QueryRow())
		if err != nil {
			return nil, err
		}
	}

	return &dbSchool, nil
}

func buildUpsertSchoolSql(school *school.School) (string, []sqlValue) {
	inserted := sqlValues{}
	if school.Id != 0 {
		inserted[School.id] = school.Id
	}

	inserted[School.name] = school.Name
	inserted[School.visible] = school.Visible

	if school.StartDate != nil {
		inserted[School.start] = school.StartDate
	}
	if school.EndDate != nil {
		inserted[School.end] = school.EndDate
	}

	if school.RegistrationOpenDate != nil {
		inserted[School.registrationStart] = school.RegistrationOpenDate
	}
	if school.RegistrationCloseDate != nil {
		inserted[School.registrationEnd] = school.RegistrationCloseDate
	}

	cols, vals := inserted.split()
	updatedCols := make([]colName, len(cols))
	for i, col := range cols {
		updatedCols[i] = fmt.Sprintf("%[1]s=EXCLUDED.%[1]s", col)
	}

	return fmt.Sprintf(
		`INSERT INTO school (%s) VALUES (%s) 
		ON CONFLICT (id) DO UPDATE 
			SET %s
		RETURNING %s`,
		cols.sqlString(), cols.sqlParams(), strings.Join(updatedCols, ", "), allSchoolColumns.sqlString(),
	), vals
}

func mapSchool(row pgx.Row) (s school.School, err error) {
	err = row.Scan(
		&s.Id,
		&s.Name,
		&s.Visible,
		&s.StartDate,
		&s.EndDate,
		&s.RegistrationOpenDate,
		&s.RegistrationCloseDate,
	)
	return
}
