package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/julhan07/fiber-service/constants"
	"github.com/julhan07/fiber-service/utils"
)

func AuthAge(c *fiber.Ctx) error {

	age, _ := strconv.Atoi(c.Query("age"))

	if age >= 18 {
		return c.Next()
	} else {
		return utils.ResponseError(c, constants.StatusBadRequest, "Age min 18 year")
	}
}
