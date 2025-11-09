package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"news-feed-system-backend/app/models"
	"news-feed-system-backend/database"
)

func CreatePost(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	post := models.Post{
		UserID:    userID,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&post).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id":        post.ID,
		"userid":    post.UserID,
		"content":   post.Content,
		"createdat": post.CreatedAt,
	})
}

func GetFeed(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var posts []models.Post
	if err := database.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch feed"})
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"posts": posts,
	})
}
