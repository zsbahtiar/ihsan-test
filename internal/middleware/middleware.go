package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/zsbahtiar/ihsan-test/internal/pkg/logger"
	"go.uber.org/zap"
	"time"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(recover.New())

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	app.Use(Logger())

}

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		fields := []zap.Field{
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
		}

		if err != nil {
			logger.Error("Request failed",
				append(fields, zap.Error(err))...,
			)
			return err
		}

		if c.Response().StatusCode() >= 400 {
			body := c.Response().Body()
			fields = append(fields, zap.String("response", string(body)))
			logger.Error("Request failed", fields...)
			return err
		}

		logger.Info("Request completed", fields...)
		return nil
	}
}
