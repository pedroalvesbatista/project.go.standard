package hcustomer

import (
	mcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/models/customer"
	mErrors "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/models/errors"
	fmts "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/pkg/fmts"
	rcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/repo/customer"
	"github.com/gofiber/fiber"
)

func (s *Server) Post(c *fiber.Ctx) {
	var input mcustomer.CustomerPost
	var Errors mErrors.Errors
	if err := c.BodyParser(&input); err != nil {
		Errors.Msg = err.Error()
		c.Status(503).JSON(Errors)
		return
	}
	input.ImpIp = c.IP()
	if err := rcustomer.Post(s.Db, input); err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//503 Serviço indisponivel
		c.Status(503).JSON(Errors)
		return
	}
	c.Status(200)

}
