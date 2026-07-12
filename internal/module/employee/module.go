package employee

import (
	"github.com/TrueFlowDev/Backend/internal/module/employee/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/employee/infrastructure/adapter"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"employee",
	fx.Provide(
		fx.Annotate(
			adapter.NewEmployeeRepository,
			fx.As(new(port.EmployeeRepository)),
		),
	),
)
