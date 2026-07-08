package organization

import (
	"github.com/TrueFlowDev/Backend/internal/module/organization/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/organization/infrastructure/adapter"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"organization",
	fx.Provide(
		fx.Annotate(
			adapter.NewOrganizationRepository,
			fx.As(new(port.OrganizationRepository)),
		),
	),
)
