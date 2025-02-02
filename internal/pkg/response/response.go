package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code   string `json:"code"`
	Remark string `json:"remark"`
}

func WriteSuccess(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func WriteError(c *fiber.Ctx, err error) error {
	if e, ok := err.(*Error); ok {
		return c.Status(e.StatusCode).JSON(Response{
			Success: false,
			Error: &ErrorInfo{
				Code:   e.Code,
				Remark: e.Remark,
			},
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Success: false,
		Error: &ErrorInfo{
			Code:   "INTERNAL_ERROR",
			Remark: "internal server error",
		},
	})
}
