package server

import (
	"fmt"
	"goly/model"
	"goly/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func redirect(ctx *fiber.Ctx) error {
	goly, err := model.FindByGolyUrl(ctx.Params("redirect"))
	golyURL := goly.Redirect

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cant find goly url " + err.Error(),
		})
	}

	goly.Clicked++
	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Print("error updating click count", err)
	}

	return ctx.Redirect(golyURL, fiber.StatusTemporaryRedirect)
}

func getGolies(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGolies()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	goly, err := model.GetGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	goly := model.Goly{}
	err := ctx.BodyParser(&goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing json " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error couldnt create goly in db " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goly model.Goly
	err := ctx.BodyParser(&goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing json" + err.Error(),
		})
	}

	// if goly.Random {
	// 	goly.Goly = utils.RandomURL(8)
	// }

	err = model.UpdateGoly(goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error updating db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing json" + err.Error(),
		})
	}

	err = model.DeleteGoly(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting from db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "delete success",
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.Get("/goly", getGolies)
	router.Get("/r/:redirect", redirect)
	router.Get("/goly/:id", getGoly)
	router.Post("/goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Delete("/goly", deleteGoly)
	// router.Get("/redirect", )

	log.Fatal(router.Listen(":3000"))
}
