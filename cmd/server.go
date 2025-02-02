package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zsbahtiar/ihsan-test/internal/middleware"
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

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	return app.Listen(addr)
}
