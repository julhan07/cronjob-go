package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julhan07/fiber-service/constants"
	"github.com/julhan07/fiber-service/utils"
)

func (p *presenter) GetUsers(c *fiber.Ctx) error {
	return utils.ResponseSuccess(c, constants.StatusOK, "ok", nil)
}

func (p *presenter) PostUser(c *fiber.Ctx) error {
	return utils.ResponseSuccess(c, constants.StatusCreated, "ok", nil)
}
