package controllers

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"bikrent/models"
	"log"
	"net/http"
	"strconv"
)

type BicycleController struct {
	DB *gorm.DB
}

func NewBicycleController(db *gorm.DB) *BicycleController {
	return &BicycleController{
		DB: db,
	}
}

func (bc *BicycleController) GetBicycles(c echo.Context) error {
	var bicycles []models.Bicycle
	if err := bc.DB.Find(&bicycles).Error; err != nil {
		log.Println("Error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve bicycles"})
	}

	return c.JSON(http.StatusOK, bicycles)
}

func (bc *BicycleController) GetBicycleByID(c echo.Context) error {
    bikeID := c.Param("id")

    var bike models.Bicycle
    if err := bc.DB.First(&bike, bikeID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Bike not found"})
    }

    return c.JSON(http.StatusOK, bike)
}

func (bc *BicycleController) CreateBicycle(c echo.Context) error {
    name := c.FormValue("name")
    bikeType := c.FormValue("type")
    seat := c.FormValue("seat")
    pricePerHourStr := c.FormValue("price_per_hour")
    pricePerHour, err := strconv.Atoi(pricePerHourStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price_per_hour value"})
    }

    bicycle := &models.Bicycle{
        Name:        name,
        Type:        bikeType,
        Seat:        seat,
        PricePerHour: pricePerHour,
    }

    if err := bc.DB.Create(bicycle).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create bicycle"})
    }

    return c.JSON(http.StatusCreated, bicycle)
}

func (bc *BicycleController) UpdateBicycle(c echo.Context) error {
    bicycleID := c.Param("id")

    var bicycle models.Bicycle
    if err := bc.DB.First(&bicycle, bicycleID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Bicycle not found"})
    }

    name := c.FormValue("name")
    bikeType := c.FormValue("type")
    seat := c.FormValue("seat")
    pricePerHourStr := c.FormValue("price_per_hour")

    if name != "" {
        bicycle.Name = name
    }
    if bikeType != "" {
        bicycle.Type = bikeType
    }
    if seat != "" {
        bicycle.Seat = seat
    }
    if pricePerHourStr != "" {
        pricePerHour, err := strconv.Atoi(pricePerHourStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid price_per_hour value"})
        }
        bicycle.PricePerHour = pricePerHour
    }

    if err := bc.DB.Save(&bicycle).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update bicycle"})
    }

    return c.JSON(http.StatusOK, bicycle)
}


func (bc *BicycleController) DeleteBicycle(c echo.Context) error {
    bicycleID := c.Param("id")

    var bicycle models.Bicycle
    if err := bc.DB.First(&bicycle, bicycleID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Bicycle not found"})
    }

    if err := bc.DB.Delete(&bicycle).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete bicycle"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Bicycle deleted successfully"})
}