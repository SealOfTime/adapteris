package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

type StageHandlers struct {
	*fiber.App
	log          log.Logger
	stageService *school.StageService
	stepService  *school.StepService
}

func NewStageHandlers(
	log log.Logger,
	stageService *school.StageService,
	stepService *school.StepService,
) *StageHandlers {
	app := &StageHandlers{
		App:          fiber.New(),
		log:          log,
		stageService: stageService,
		stepService:  stepService,
	}
	stage := app.Group("/:stageId")
	{
		stage.Put("/name", app.RenameStage)
		stage.Put("/description", app.ChangeStageDescription)
		stage.Put("/dates", app.ScheduleStage)
		stage.Post("/steps", app.AddStep)
	}

	return app
}

type Stage struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start"`
	EndDate     *time.Time `json:"end"`
	Steps       []Step     `json:"steps"`
}

func (h *StageHandlers) RenameStage(c *fiber.Ctx) error {
	//Path parameter.
	var (
		stageId int64
	)
	type RequestBody struct {
		NewName string `json:"newName"`
	}

	var err error
	if stageId, err = strconv.ParseInt(c.Params("stageId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad stageId",
		)
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx := c.UserContext()

	if err := h.stageService.Rename(ctx, school.StageRenameRequest{
		StageId: stageId,
		NewName: body.NewName,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *StageHandlers) ChangeStageDescription(c *fiber.Ctx) error {
	//Path parameter.
	var (
		stageId int64
	)
	type RequestBody struct {
		NewDescription string `json:"newDescription"`
	}

	var err error
	if stageId, err = strconv.ParseInt(c.Params("stageId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad stageId",
		)
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx := c.UserContext()
	if err := h.stageService.ChangeDescription(ctx, school.StageChangeDescriptionRequest{
		StageId:        stageId,
		NewDescription: body.NewDescription,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *StageHandlers) ScheduleStage(c *fiber.Ctx) error {
	//Path parameter.
	var (
		stageId int64
	)
	type RequestBody struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}

	var err error
	if stageId, err = strconv.ParseInt(c.Params("stageId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad stageId",
		)
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx := c.UserContext()
	if err := h.stageService.ScheduleStage(ctx, school.SchedulStageRequest{
		StageId: stageId,
		Start:   body.Start,
		End:     body.End,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *StageHandlers) AddStep(c *fiber.Ctx) error {
	//Path param
	var (
		stageId int64
	)
	type Response struct {
		Step Step `json:"step"`
	}

	var (
		rawStageId string
		err        error
	)
	if rawStageId = c.Params("stageId"); rawStageId == "my" {

	}

	if stageId, err = strconv.ParseInt(rawStageId, 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad stageId",
		)
	}

	ctx := c.UserContext()

	step, err := h.stepService.AddStep(ctx, school.AddStepRequest{
		StageId: stageId,
	})
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.JSON(Response{
		Step: domainStepToDto(*step),
	})
}

func domainStageToDto(stage school.Stage) Stage {
	res := Stage{
		Id:          stage.Id,
		Name:        stage.Name,
		Description: stage.Description,
		StartDate:   stage.StartDate,
		EndDate:     stage.EndDate,
	}
	if len(stage.Steps) != 0 {
		res.Steps = make([]Step, len(stage.Steps))
		for i, step := range stage.Steps {
			res.Steps[i] = domainStepToDto(step)
		}
	}
	return res
}
