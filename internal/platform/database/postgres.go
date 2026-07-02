package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TrueFlowDev/Backend/internal/platform/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	host     string
	port     int
	user     string
	password string
	dbName   string

	driver *gorm.DB
	sql    *sql.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	return &Postgres{
		host:     cfg.DB.PostgresHost,
		port:     cfg.DB.PostgresPort,
		user:     cfg.DB.PostgresUser,
		password: cfg.DB.PostgresPassword,
		dbName:   cfg.DB.PostgresDB,
	}
}

func (p *Postgres) Connect() error {
	db, err := gorm.Open(
		postgres.Open(p.dsn()),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	p.driver = db
	p.sql = sqlDB

	return nil
}

func (p *Postgres) Close() error {
	return p.sql.Close()
}

func (p *Postgres) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.host,
		p.port,
		p.user,
		p.password,
		p.dbName,
	)
}
