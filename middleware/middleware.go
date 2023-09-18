package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bahati-hakizimana/blogue-backend/util"
)

func IsAuthanticate(c *fiber.Ctx) error{
	cookie:=c.Cookies("jwt")

	if _, err:=util.Parsejwt(cookie);err != nil{
	      c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message":"unauthanticated",
		})
	}
	return c.Next()
}