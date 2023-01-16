package handlers

import (
	"strconv"
	"tester-go-docker/database"

	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"Hello World nishant",
		})
	}

func CreateFact(c *fiber.Ctx) error {
	var fact database.Fact

	c.Accepts("applications/json")
	err:= c.BodyParser(&fact)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"error parsing Json"+err.Error(),
		})
	}

	err = database.CreateFact(fact)
	if err!= nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create Fact in db" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func GetAllFacts(app *fiber.Ctx) error {
	facts, err := database.GetAllFacts()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all facts" + err.Error(),
		})
	}
	return app.Status(fiber.StatusOK).JSON(facts)
}

func DeleteFact(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = database.DeleteFact(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "fact deleted.",
	})
}

func GetFact(app *fiber.Ctx) error {
	id, err := strconv.ParseUint( app.Params("id"), 10, 64)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id" + err.Error(),
		})
	}

	fact, err := database.GetFact(id)
	if err!= nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fact not found" + err.Error(),
		})
	}
	return app.Status(fiber.StatusOK).JSON(fact)
}