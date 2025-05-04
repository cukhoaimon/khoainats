package main

import (
	"log"
	"time"

	"github.com/cukhoaimon/khoainats/internal/repository"
	"github.com/cukhoaimon/khoainats/internal/repository/model"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	ById(id uuid.UUID) (gen.T, error)

	// UPDATE @@table
	// SET deleted_at = now(), deleted_by = @by
	// WHERE id = @id
	DeleteById(id, by uuid.UUID) (gen.T, error)
}

func main() {
	// TODO: fetch from Vault
	sqlDb, err := repository.NewDatabase(repository.DatabaseConfig{
		Host:            "localhost",
		Port:            "5432",
		User:            "local",
		Password:        "local",
		Dbname:          "khoainats",
		ConnMaxIdleTime: 10 * time.Minute,
		ConnMaxLifetime: 10 * time.Minute,
		MaxIdleConns:    10,
		MaxOpenConns:    10,
	})
	if err != nil {
		log.Fatal("Error when create db")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error when create gorm.db instance")
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	g.ApplyInterface(
		func(Querier) {},
		model.AccessToken{},
		model.LoginEvent{},
		model.Organization{},
		model.Principal{},
		model.PrincipalAttribute{},
	)
	g.Execute()
}
