package controllers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/k0msak007/go-todo-app/database"
	"github.com/k0msak007/go-todo-app/models"
	"github.com/k0msak007/go-todo-app/request"
)

func CreateTodo(c *fiber.Ctx) error {
	todoReq := request.TodoCreateRequest{}
	// todoReq := request.TodoCreateRequest{}

	if err := c.BodyParser(&todoReq); err != nil {
		fmt.Printf("todo request is failed %v \n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "todo request is failed",
			"error":   err.Error(),
		})
	}

	// Validation Request Data
	validate := validator.New()
	if err := validate.Struct(&todoReq); err != nil {
		fmt.Println(&todoReq)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "some date is not valid",
			"error":   err.Error(),
		})
	}

	fmt.Println(&todoReq)

	todo := models.Todo{}
	todo.Name = todoReq.Name
	todo.IsComplete = todoReq.IsComplete
	if todoReq.Note != "" {
		todo.Note = todoReq.Note
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		log.Println("todo.controller.go => CreateTodo :: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "todo created successfully",
		"data":    &todo,
	})
}

func GetAllTodo(c *fiber.Ctx) error {
	todos := make([]models.Todo, 0)

	if err := database.DB.Find(&todos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data transmited",
		"data":    &todos,
	})
}

func GetTodoByID(c *fiber.Ctx) error {
	todoId := c.Params("id")
	if todoId == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Param is required",
		})
	}
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "todo not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "data tranmitted",
		"data":    &todo,
	})
}

func UpdateTodoByID(c *fiber.Ctx) error {
	todoReq := request.TodoUpdateRequest{}
	// todoReq := request.TodoCreateRequest{}

	if err := c.BodyParser(&todoReq); err != nil {
		fmt.Printf("todo request is failed %v \n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "todo request is failed",
			"error":   err.Error(),
		})
	}

	// Validation Request Data
	validate := validator.New()
	if err := validate.Struct(&todoReq); err != nil {
		fmt.Println(&todoReq)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "some date is not valid",
			"error":   err.Error(),
		})
	}

	todoId := c.Params("id")
	if todoId == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Param is required",
		})
	}
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "todo not found",
			"error":   err.Error(),
		})
	}

	todo.Name = todoReq.Name
	todo.Note = todoReq.Note
	todo.IsComplete = todoReq.IsComplete

	if err := database.DB.Save(&todo).Error; err != nil {
		log.Println("todo.controller.go => CreateTodo :: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "todo created successfully",
		"data":    &todo,
	})
}

func DeleteTodoByID(c *fiber.Ctx) error {
	todoId := c.Params("id")
	if todoId == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Param is required",
		})
	}
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "todo not found",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Delete(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "todo deleted",
	})
}
