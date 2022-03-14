package app

import (
	"fmt"

	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/domain/user"
	"github.com/sealoftime/adapteris/integration/vk"
)

//Services stores application-level services i.e. those dependent on the concrete protocols etc
type Services struct {
	Integration struct {
		Vk *vk.Service
	}
	Auth   *user.AuthService
	User   *user.Service
	School *school.Service
	Stage  *school.StageService
	Step   *school.StepService
	Event  *school.EventService
}

func (a *App) initServices() {
	a.Service.Auth = user.NewAuthService(
		a.Storage.accounts,
	)
	a.Service.Integration.Vk = vk.New(
		a.Service.Auth,
		a.Config.Vk.ClientId,
		a.Config.Vk.Secret,
		fmt.Sprintf("%s/api/auth/vk/callback", a.Config.HostURL),
	)
	a.Service.School = school.NewService(
		a.Storage.schools,
	)
	a.Service.Stage = school.NewStageService(
		a.Log,
		a.Storage.schools,
	)
	a.Service.Step = school.NewStepsService(
		a.Log,
		a.Storage.stages,
		a.Storage.steps,
	)
	a.Service.Event = school.NewEventService(
		a.Log,
		a.Storage.steps,
		a.Storage.events,
	)
}
