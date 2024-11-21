package main

import (
	"log"
	"situs-forum/internal/configs"
	"situs-forum/internal/handlers/memberships"
	"situs-forum/internal/handlers/posts"
	membershipRepo "situs-forum/internal/repository/memberships"
	postRepo "situs-forum/internal/repository/posts"
	membershipSvc "situs-forum/internal/service/memberships"
	postSvc "situs-forum/internal/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	merbershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, merbershipService)
	postHandler := posts.NewHandler(r, postService)

	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
