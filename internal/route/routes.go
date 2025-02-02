package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zsbahtiar/ihsan-test/internal/handler"
)

func Setup(app *fiber.App, accountHandler handler.AccountHandler) {
	app.Post("/daftar", accountHandler.Register)
	app.Post("/tabung", accountHandler.Deposit)
	app.Post("/tarik", accountHandler.Withdraw)

}
