package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/school"
)

type SchoolHandlers struct {
	*fiber.App
	schoolService *school.Service
	stageService  *school.StageService
}

func NewSchoolHandlers(
	schoolService *school.Service,
	stageService *school.StageService,
) *SchoolHandlers {
	app := &SchoolHandlers{
		App:           fiber.New(),
		schoolService: schoolService,
		stageService:  stageService,
	}
	app.Post("/", app.CreateSchool)
	school := app.Group("/:schoolId")
	{
		school.Get("/", app.GetSchool)
		school.Put("/name", app.RenameSchool)
		school.Put("/dates", app.ScheduleSchool)
		school.Put("/registration", app.ScheduleRegistration)
		school.Post("/stages", app.AddStage)
	}
	return app
}

type School struct {
	Id                    int64      `json:"id"`
	Name                  string     `json:"name"`
	Visible               bool       `json:"visible"`
	StartDate             *time.Time `json:"start"`
	EndDate               *time.Time `json:"end"`
	RegistrationOpenDate  *time.Time `json:"registerStart"`
	RegistrationCloseDate *time.Time `json:"registerEnd"`
	Stages                []Stage    `json:"stages,omitempty"`
}

func (h *SchoolHandlers) GetSchool(c *fiber.Ctx) error {
	type Response struct {
		School School `json:"school"`
	}

	schoolId, err := strconv.ParseInt(c.Params("schoolId"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad schoolId",
		)
	}

	ctx := c.UserContext()
	sch, err := h.schoolService.GetSchool(ctx, schoolId)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}
	return c.JSON(
		Response{
			School: domainSchoolToDto(*sch),
		},
	)
}

func (h *SchoolHandlers) CreateSchool(c *fiber.Ctx) error {
	type Response struct {
		School School `json:"school"`
	}
	var req school.CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed body: %+v", err),
		)
	}
	ctx := c.UserContext()

	sch, err := h.schoolService.Create(ctx, req)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}
	return c.JSON(
		Response{
			School: domainSchoolToDto(*sch),
		},
	)
}

func (h *SchoolHandlers) RenameSchool(c *fiber.Ctx) error {
	schoolId, err := strconv.ParseInt(c.Params("schoolId"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad schoolId",
		)
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("malformed body: %+v", err),
		)
	}

	ctx := c.UserContext()
	if err := h.schoolService.Rename(ctx, school.RenameRequest{
		SchoolId: schoolId,
		NewName:  req.Name,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *SchoolHandlers) ScheduleSchool(c *fiber.Ctx) error {
	schoolId, err := strconv.ParseInt(c.Params("schoolId"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad schoolId",
		)
	}

	var req struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("malformed body: %+v", err),
		)
	}

	ctx := c.UserContext()
	if err := h.schoolService.ScheduleSchool(ctx, school.ScheduleSchoolReq{
		SchoolId:  schoolId,
		StartDate: req.Start,
		EndDate:   req.End,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *SchoolHandlers) ScheduleRegistration(c *fiber.Ctx) error {
	schoolId, err := strconv.ParseInt(c.Params("schoolId"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad schoolId",
		)
	}

	var req struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("malformed body: %+v", err),
		)
	}

	ctx := c.UserContext()
	if err := h.schoolService.ScheduleRegistration(ctx, school.ScheduleRegistrationReq{
		SchoolId:  schoolId,
		StartDate: req.Start,
		EndDate:   req.End,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *SchoolHandlers) AddStage(c *fiber.Ctx) error {
	type Response struct {
		Stage Stage `json:"stage"`
	}

	schoolId, err := strconv.ParseInt(c.Params("schoolId"), 10, 64)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad schoolId",
		)
	}

	var req struct {
		Stage Stage `json:"stage"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("malformed body: %+v", err),
		)
	}

	ctx := c.UserContext()
	st, err := h.stageService.AddStage(ctx, school.AddStageRequest{
		SchoolId:    schoolId,
		Name:        req.Stage.Name,
		Description: req.Stage.Description,
		Start:       req.Stage.StartDate,
		End:         req.Stage.EndDate,
	})
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.JSON(Response{
		Stage: domainStageToDto(*st),
	})
}

func domainSchoolToDto(sch school.School) School {
	schoolDto := School{
		Id:                    sch.Id,
		Name:                  sch.Name,
		Visible:               sch.Visible,
		StartDate:             sch.StartDate,
		EndDate:               sch.EndDate,
		RegistrationOpenDate:  sch.RegistrationOpenDate,
		RegistrationCloseDate: sch.RegistrationCloseDate,
	}
	if len(sch.Stages) != 0 {
		schoolDto.Stages = make([]Stage, len(sch.Stages))
		for i, stage := range sch.Stages {
			schoolDto.Stages[i] = domainStageToDto(stage)
		}
	}
	return schoolDto
}
