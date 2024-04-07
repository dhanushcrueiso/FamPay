package handlers

import (
	"fmt"
	"strconv"

	"FAMPAY/internal/dtos"
	"FAMPAY/internal/services"

	"github.com/gofiber/fiber"
)

func GetYoutubeDataPaginated(c *fiber.Ctx) {
	fmt.Println("here:")

	qu := c.Query("q")
	page, err := strconv.Atoi(c.Query("pg"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var DataFilter = dtos.DataFilter{
		Page: page,
		Q:    qu,
	}
	if err := c.BodyParser(&DataFilter); err != nil {
		return
	}
	fmt.Println("checking the parser", DataFilter)

	res, err := services.GetDataPaginated(DataFilter)
	if err != nil {
		return
	}
	c.JSON(res)
}
