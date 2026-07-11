package main

import (
	"github.com/TrueFlowDev/Backend/internal/module/authentication"
	"github.com/TrueFlowDev/Backend/internal/module/employee"
	"github.com/TrueFlowDev/Backend/internal/module/notification"
	"github.com/TrueFlowDev/Backend/internal/module/organization"
	"github.com/TrueFlowDev/Backend/internal/module/user"
	"github.com/TrueFlowDev/Backend/internal/platform"
	"github.com/TrueFlowDev/Backend/internal/shared"
	"go.uber.org/fx"
)

//	@title			TrueFlow Api Document
//	@description	TrueFlow is a modern human resource management (HRM) platform designed to integrate organizational processes, manage teams, control access, and increase organizational productivity.

//	@contact.name	API Support
//	@contact.url	https://trueflow.ir/support
//	@contact.email	support@trueflow.ir

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Enter the JWT token with the `Bearer ` prefix
func main() {
	fx.
		New(
			shared.Module,
			platform.Module,
			user.Module,
			authentication.Module,
			notification.Module,
			organization.Module,
			employee.Module,
		).
		Run()
}
