package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
    "bikrent/models"
)

func TestRentalController_CreateRental(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    rc := NewRentalController(db)

    // Ensure a clean database by deleting all rentals (if necessary)
    db.Exec("DELETE FROM rentals")

    // Create a test rental data
    formData := "bicycle_id=1&user_id=1&rental_start_time=2023-10-31 10:00:00&rental_end_time=2023-10-31 12:00:00"
    req := httptest.NewRequest(http.MethodPost, "/rentals", strings.NewReader(formData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

    // Create a recorder for the response
    rec := httptest.NewRecorder()

    // Create an Echo context and bind the controller method
    c := e.NewContext(req, rec)
    err := rc.CreateRental(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, rec.Code)
    // You can also assert the response body to check the created rental data.
}

func TestRentalController_GetRental(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    rc := NewRentalController(db)

    // Create a test rental
    testRental := models.Rental{
        BicycleID:      1,
        UserID:         1,
        RentalStartTime: "2023-10-31 10:00:00", // Use the appropriate date and time string
        RentalEndTime:   "2023-10-31 12:00:00", // Use the appropriate date and time string
        TotalPrice:     10.0, // Set the appropriate total price
    }
    db.Create(&testRental)

    // Define a test request to get a rental by ID
    req := httptest.NewRequest(http.MethodGet, "/rentals/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    err := rc.GetRental(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    // You can also assert the response body to check the rental data.
}
