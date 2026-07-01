package main

import (
	"github.com/TrueFlowDev/Backend/internal/module/user"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			user.Module,
		).
		Run()
}
