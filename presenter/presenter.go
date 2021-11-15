package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julhan07/fiber-service/constants"
	"github.com/julhan07/fiber-service/middleware"
	"github.com/julhan07/fiber-service/service"
	"github.com/julhan07/fiber-service/utils"
)

type presenter struct {
	route   *fiber.App
	service *service.ServiceIndex
}

func NewPresenterIndex(r *fiber.App, service *service.ServiceIndex) {

	p := &presenter{
		route:   r,
		service: service,
	}
	p.routes()
}

func (p *presenter) routes() {

	p.route.Get("/", p.Welcome)
	p.route.Post("/user/register", middleware.AuthAge, p.PostUser)
	p.route.Get("/user/list", middleware.AuthAge, p.GetUsers)

}

func (p *presenter) Welcome(c *fiber.Ctx) error {
	return utils.ResponseSuccess(c, constants.StatusOK, "ok", nil) // => 👴 julhan is 28 year
}
