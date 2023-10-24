package controllers

import (
    "github.com/labstack/echo"
	// "golang.org/x/crypto/bcrypt"
    // "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
    "bikrent/models"
	"log"
	"time"
	"strconv"
)

type RentalController struct {
	DB *gorm.DB
}

func NewRentalController(db *gorm.DB) *RentalController {
	return &RentalController{
		DB: db,
	}
}

func (rc *RentalController) CreateRental(c echo.Context) error {
    rental := new(models.Rental)

    bicycleIDStr := c.FormValue("bicycle_id")
    userIDStr := c.FormValue("user_id")
    rentalStartTimeStr := c.FormValue("rental_start_time")
    rentalEndTimeStr := c.FormValue("rental_end_time")

    bicycleID, err := strconv.Atoi(bicycleIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid bicycle_id format"})
    }
    rental.BicycleID = uint(bicycleID)

    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id format"})
    }
    rental.UserID = uint(userID)

    startTime, err := time.Parse("2006-01-02 15:04:05", rentalStartTimeStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid rental_start_time format"})
    }
    endTime, err := time.Parse("2006-01-02 15:04:05", rentalEndTimeStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid rental_end_time format"})
    }

    duration := endTime.Sub(startTime)
    hours := duration.Hours()

    // Fetch the bicycle's price per hour from the database
    var bicycle models.Bicycle
    if err := rc.DB.First(&bicycle, rental.BicycleID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Bicycle not found"})
    }

    totalPrice := hours * float64(bicycle.PricePerHour)
    rental.RentalStartTime = startTime.Format(time.RFC3339)
    rental.RentalEndTime = endTime.Format(time.RFC3339)
    rental.TotalPrice = totalPrice

    if err := rc.DB.Create(rental).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create rental"})
    }

    return c.JSON(http.StatusCreated, rental)
}



func (rc *RentalController) GetRental(c echo.Context) error {
    rentalID := c.Param("id")

    var rental models.Rental

    if err := rc.DB.Preload("User").Preload("Bicycle").First(&rental, rentalID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Rental not found"})
    }

    return c.JSON(http.StatusOK, rental)
}

