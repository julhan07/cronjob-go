package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/julhan07/fiber-service/presenter"
	"github.com/julhan07/fiber-service/repo"
	"github.com/julhan07/fiber-service/service"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Static("/assets", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	repo := repo.NewRepoIndex()
	service := service.NewServiceIndex(repo)
	presenter.NewPresenterIndex(app, &service)

	c := cron.New()

	c.AddFunc("@every 0h1m", service.CronJob)
	go func() {
		c.Start()
	}()

	app.Listen(":3000")
}
