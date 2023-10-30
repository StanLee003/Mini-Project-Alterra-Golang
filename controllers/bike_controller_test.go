package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
	"gorm.io/gorm"
    "testing"
    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
    "bikrent/models"
)

func TestBicycleController_GetBicycles(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    bc := NewBicycleController(db)

    // Define a test request to get all bicycles
    req := httptest.NewRequest(http.MethodGet, "/bicycles", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := bc.GetBicycles(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    // You can also assert the response body to check the data returned.
}

func TestBicycleController_GetBicycleByID(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    bc := NewBicycleController(db)

    // Create a test bicycle
    testBike := models.Bicycle{
        Name:        "TestBike",
        Type:        "Mountain",
        Seat:        "1",
        PricePerHour: 10,
    }
    db.Create(&testBike)

    // Define a test request to get a bicycle by ID
    req := httptest.NewRequest(http.MethodGet, "/bicycles/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    err := bc.GetBicycleByID(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    // You can also assert the response body to check the data returned.
}

func TestBicycleController_CreateBicycle(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    bc := NewBicycleController(db)

    // Ensure a clean database by deleting all bicycles (if necessary)
    db.Exec("DELETE FROM bicycles")
    db.Exec("ALTER TABLE bicycles AUTO_INCREMENT = 1")

    // Define a test request to create a bicycle
    formData := "name=TestBike&type=Mountain&seat=1&price_per_hour=10"
    req := httptest.NewRequest(http.MethodPost, "/bicycles", strings.NewReader(formData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

    // Create a recorder for the response
    rec := httptest.NewRecorder()

    // Create an Echo context and bind the controller method
    c := e.NewContext(req, rec)
    err := bc.CreateBicycle(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, rec.Code)
    // Check if the bicycle has been created in the database
    var bicycle models.Bicycle
    if err := db.Where("name = ?", "TestBike").First(&bicycle).Error; err != nil {
        t.Errorf("Expected bicycle to be created with name 'TestBike', but got an error: %v", err)
    }
}

func TestBicycleController_UpdateBicycle(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    bc := NewBicycleController(db)

    // Create a test bicycle
    testBike := models.Bicycle{
        Name:        "TestBike",
        Type:        "Mountain",
        Seat:        "1",
        PricePerHour: 10,
    }
    db.Create(&testBike)

    // Define a test request to update a bicycle by ID
    req := httptest.NewRequest(http.MethodPut, "/bicycles/1", strings.NewReader("name=UpdatedBike"))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    err := bc.UpdateBicycle(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    // You can also assert the response body to check the updated bicycle data.
}

func TestBicycleController_DeleteBicycle(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    bc := NewBicycleController(db)

    // Create a test bicycle
    testBike := models.Bicycle{
        Name:        "TestBike",
        Type:        "Mountain",
        Seat:        "1",
        PricePerHour: 10,
    }
    db.Create(&testBike)

    // Define a test request to delete a bicycle by ID
    req := httptest.NewRequest(http.MethodDelete, "/bicycles/1", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("1")

    err := bc.DeleteBicycle(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    // Check if the bicycle has been deleted from the database
    var deletedBike models.Bicycle
    if err := db.Where("name = ?", "TestBike").First(&deletedBike).Error; err != gorm.ErrRecordNotFound {
        t.Errorf("Expected bicycle to be deleted with name 'TestBike', but it still exists")
    }
}
