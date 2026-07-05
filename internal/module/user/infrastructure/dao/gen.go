//go:build ignore

package main

import (
	"fmt"

	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.PostgresHost,
		cfg.DB.PostgresPort,
		cfg.DB.PostgresUser,
		cfg.DB.PostgresPassword,
		cfg.DB.PostgresDB,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:        "internal/module/user/infrastructure/dao",
		ModelPkgPath:   "internal/module/user/infrastructure/model",
		FieldNullable:  true,
		FieldCoverable: true,
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateModel("users"))
	g.Execute()
}
