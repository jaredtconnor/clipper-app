package clippingHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jaredtconnor/clipper-app/database"
	"github.com/jaredtconnor/clipper-app/internal/model"
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
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": clipping})

}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var clipping model.Clipping

	id := c.Params("clippingId")

	db.Find(&clipping, "id = ?", id)

	if clipping.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clipping present", "data": nil})

	}

	return c.JSON(fiber.Map{"status": "success", "messgae": "Note found", "data": clipping})
}

func UpdateClipping(c *fiber.Ctx) {

}

func DeleteClipping(c *fiber.Ctx) {

	db := database.DB
	var clipping model.Clipping

	id := c.Params("clippingId")

	db.Find(&clipping, "id = ?", id)

	if clipping.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clipping present", "data": nil})
	}

	err := db.Delete(&clipping, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Clipping"})

}
