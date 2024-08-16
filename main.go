package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/swagger"
    _ "bc4m-case/docs"
    "time"
    "fmt"
)

type HealthStatus struct {
    Status string `json:"status"`
    Time   string `json:"time"`
}

// @title Fiber API
// @version 1.0
// @description This is a simple API for BC4M.
// @host localhost:8080
// @BasePath /
func main() {
    app := fiber.New()

    app.Use(recover.New())

    app.Get("/swagger/*", swagger.HandlerDefault)

    defineRoutes(app)

    app.Listen(":8080")
}

func defineRoutes(app *fiber.App) {
    app.Get("/", getMessage)
    app.Get("/health", getHealthStatus)
    app.Post("/", postData)
}

// getMessage handles GET requests to the root endpoint.
// @Summary Get message
// @Description Get a static message.
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func getMessage(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "msg": "BC4M",
    })
}

// getHealthStatus handles GET requests to the /health endpoint.
// @Summary Get health status
// @Description Get the health status of the application.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthStatus
// @Router /health [get]
func getHealthStatus(c *fiber.Ctx) error {
    status := HealthStatus{
        Status: "OK",
        Time:   time.Now().Format(time.RFC3339),
    }
    return c.Status(fiber.StatusOK).JSON(status)
}

// postData handles POST requests to the root endpoint.
// @Summary Echo data
// @Description Echo the data posted in the request body.
// @Tags root
// @Accept json
// @Produce json
// @Param data body map[string]interface{} true "Data"
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func postData(c *fiber.Ctx) error {
    return c.Send(c.Body())
}

