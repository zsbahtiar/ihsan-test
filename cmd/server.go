package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zsbahtiar/ihsan-test/internal/core/module"
	"github.com/zsbahtiar/ihsan-test/internal/core/repository"
	"github.com/zsbahtiar/ihsan-test/internal/handler"
	"github.com/zsbahtiar/ihsan-test/internal/middleware"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/database"
	"github.com/zsbahtiar/ihsan-test/internal/route"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Long:  `Start the HTTP server account-service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer()
	},
}

func runServer() error {
	app := fiber.New(fiber.Config{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	})
	middleware.SetupMiddleware(app)
	db := database.NewPostgres(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	defer db.Close()
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := module.NewAccountUsecase(accountRepo)
	accountHandler := handler.NewAccountHandler(accountUsecase)

	route.Setup(app, accountHandler)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	return app.Listen(addr)
}
