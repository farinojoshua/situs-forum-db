package main

import (
	"log"
	"situs-forum/internal/configs"
	"situs-forum/internal/handlers/memberships"
	membershipRepo "situs-forum/internal/repository/memberships"
	membershipSvc "situs-forum/internal/service/memberships"
	"situs-forum/pkg/internalsql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("gagal inisiasi db", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	merbershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(r, merbershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
