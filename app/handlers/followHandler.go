package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"news-feed-system-backend/app/models"
	"news-feed-system-backend/database"
)

func Follow(c *fiber.Ctx) error {
	followerID := c.Locals("user_id").(uint)
	followeeIDParam := c.Params("userid")

	followeeID64, err := strconv.ParseUint(followeeIDParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	followeeID := uint(followeeID64)

	if followerID == followeeID {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "You cannot follow yourself"})
	}

	follow := models.Follow{FollowerID: followerID, FolloweeID: followeeID}
	if err := database.DB.Create(&follow).Error; err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Already following"})
	}

	return c.JSON(fiber.Map{"message": "You are now following user " + followeeIDParam})
}

func Unfollow(c *fiber.Ctx) error {
	followerID := c.Locals("user_id").(uint)
	followeeIDParam := c.Params("userid")

	followeeID64, err := strconv.ParseUint(followeeIDParam, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	followeeID := uint(followeeID64)

	if err := database.DB.Where("follower_id = ? AND followee_id = ?", followerID, followeeID).Delete(&models.Follow{}).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unfollow"})
	}

	return c.JSON(fiber.Map{"message": "You unfollowed user " + followeeIDParam})
}

func GetFollowing(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or missing user ID"})
	}

	var followingIDs []uint
	if err := database.DB.Model(&models.Follow{}).
		Where("follower_id = ?", userID).
		Pluck("followee_id", &followingIDs).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch following"})
	}

	return c.JSON(fiber.Map{
		"following": followingIDs,
	})
}
