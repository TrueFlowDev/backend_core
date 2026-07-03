package main

import (
	"github.com/TrueFlowDev/Backend/internal/shared/infrastructure/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/shared/infrastructure/dao",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithGeneric,
	})

	g.ApplyBasic(model.User{})

	g.Execute()
}
