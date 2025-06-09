package routes

import (
	"github.com/FilledEther20/URL_Shortner/database"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(ctx *fiber.Ctx) error {
	url := ctx.Params("url")
	redis := database.CreateRedisClient(0)
	defer redis.Close()

	short, err := redis.Get(database.Ctx, url).Result()

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in databaset",
		})
	}

	redisInr := database.CreateRedisClient(1)
	defer redisInr.Close()

	_ = redisInr.Incr(database.Ctx, "counter")
	return ctx.Redirect(short, 301)
}
