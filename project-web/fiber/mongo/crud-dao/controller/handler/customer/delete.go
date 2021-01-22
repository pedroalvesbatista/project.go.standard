package hcustomer

import (
	fmts "github.com/project.go.standard/project-web/fiber/crud.mongoa/internal/fmts"
	mErrors "github.com/project.go.standard/project-web/fiber/crud.mongoa/models/errors"
	rcustomer "github.com/project.go.standard/project-web/fiber/crud.mongoa/repo/customer"
	"github.com/gofiber/fiber"
)

func Delete(c *fiber.Ctx) {
	var Errors mErrors.Errors
	uuid := c.Params("uuid")
	if uuid == "" {
		Errors.Msg = "Error: Nenhum uuid informado "
		//400=Bad request
		c.Status(400).JSON(Errors)
		return
	}

	if err := rcustomer.Delete(uuid); err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//404= not found
		c.Status(404).JSON(Errors)
		return
	}
	c.Status(200)
	return
}
