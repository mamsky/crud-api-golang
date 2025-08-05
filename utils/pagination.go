package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Paginate(c *fiber.Ctx) (limit int, offset int) {
	limitStr := c.Query("limit", "10") // default limit = 10
	pageStr := c.Query("page", "1")    // default page = 1

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return 0, 0 // akan divalidasi di controller
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	offset = (page - 1) * limit
	return limit, offset
}
