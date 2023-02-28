package clippingHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jaredtconnor/clipper-app/database"
	"github.com/jaredtconnor/clipper-app/internals/model"
)

func GetClippings(c *fiber.Ctx) error {
	db := database.DB

	var clippings []model.Clipping

	// Search all clippings in the database
	db.Find(&clippings)

	if len(clippings) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clippings foud", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "messgae": "Clippings found", "data": clippings})

}

func CreateClipping(c *fiber.Ctx) error {

	db := database.DB
	clipping := new(model.Clipping)

	// Store body in the note and return error if found
	err := c.BodyParser(clipping)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review input", "data": err})
	}

	clipping.ID = uuid.New()

	// Create clipping and return error if found
	err = db.Create(&clipping).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create clipping", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Clipping", "data": clipping})

}

func GetClipping(c *fiber.Ctx) error {

	db := database.DB
	var clipping model.Clipping

	id := c.Params("clippingId")

	db.Find(&clipping, "id = ?", id)

	if clipping.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clipping present", "data": nil})

	}

	return c.JSON(fiber.Map{"status": "success", "messgae": "Clipping found", "data": clipping})
}

func UpdateClipping(c *fiber.Ctx) error {

	type updateClipping struct {
		Title    string `json:"title"`
		URL      string `json:"url"`
		Contents string `json:"description"`
	}

	db := database.DB

	var clipping model.Clipping

	// Read clipping ID
	id := c.Params("clippingId")

	// Find the clipping in database
	db.Find(&clipping, "id = ?", id)

	// If no clipping present return error
	if clipping.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clipping present", "data": nil})
	}

	var updateClippingData updateClipping
	err := c.BodyParser(&updateClippingData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	clipping.Title = updateClippingData.Title
	clipping.URL = updateClippingData.URL
	clipping.Contents = updateClippingData.Contents

	db.Save(&clipping)

	return c.JSON(fiber.Map{"status": "success", "messgae": "Note found", "data": clipping})

}

func DeleteClipping(c *fiber.Ctx) error {

	db := database.DB
	var clipping model.Clipping

	id := c.Params("clippingId")

	db.Find(&clipping, "id = ?", id)

	if clipping.ID == uuid.Nil {
		return c.JSON(fiber.Map{"status": "error", "message": "No clipping present", "data": nil})
	}

	err := db.Delete(&clipping, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Clipping"})

}
