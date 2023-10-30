package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"bikrent/models"
)

func setupTestDB() *gorm.DB {
    // Define your MySQL database connection parameters
    dsn := "root:@Kotamedan3@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

    // Open a MySQL database connection using GORM
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }

    // AutoMigrate your database tables as needed
    db.AutoMigrate(&models.User{}, &models.Rental{}) // Make sure to use the correct model

    return db
}

func TestUserController_CreateUser(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    uc := NewUserController(db)

    // Ensure a clean database by deleting all users (if necessary)
    db.Exec("DELETE FROM users")
    db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

    // Define a test request with user data and a unique username
    formData := "username=uniqueuser&password=testpassword"
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(formData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

    // Create a recorder for the response
    rec := httptest.NewRecorder()

    // Create an Echo context and bind the controller method
    c := e.NewContext(req, rec)
    err := uc.CreateUser(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, rec.Code)

    // Check if the user with the username "uniqueuser" exists in the database
    var user models.User
    if err := db.Where("username = ?", "uniqueuser").First(&user).Error; err != nil {
        t.Errorf("Expected user with username 'uniqueuser' to exist, but got an error: %v", err)
    }
}

func TestUserController_Login(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    uc := NewUserController(db)

    // Ensure a clean database by deleting all users (if necessary)
    db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

    // Define a test request with an invalid username
    formData := "username=nonexistentuser&password=testpassword"
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(formData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

    // Create a recorder for the response
    rec := httptest.NewRecorder()

    // Create an Echo context and bind the controller method
    c := e.NewContext(req, rec)
    err := uc.Login(c)
	if err != nil {
		// Handle the error here
		log.Println("Error:", err)
	}
    // Assert the response
    assert.Equal(t, http.StatusUnauthorized, rec.Code)
}


// Add similar tests for other controller methods (GetUsers, GetUserByID, UpdateUser, DeleteUser) following the above patterns.

func TestUserController_GetUserByID(t *testing.T) {
	// Create a test Echo instance and a test DB
	e := echo.New()
	db := setupTestDB()
	uc := NewUserController(db)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Role:     0,
	}
	db.Create(&testUser)

	// Define a test request to get a user by ID
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := uc.GetUserByID(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	// You can also assert the response body to check the data returned.
}

func TestUserController_UpdateUser(t *testing.T) {
	// Create a test Echo instance and a test DB
	e := echo.New()
	db := setupTestDB()
	uc := NewUserController(db)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Role:     0,
	}
	db.Create(&testUser)

	// Define a test request to update a user by ID
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader("username=updateduser"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := uc.UpdateUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	// You can also assert the response body to check the updated user data.
}

func TestUserController_DeleteUser(t *testing.T) {
	// Create a test Echo instance and a test DB
	e := echo.New()
	db := setupTestDB()
	uc := NewUserController(db)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Role:     0,
	}
	db.Create(&testUser)

	// Define a test request to delete a user by ID
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := uc.DeleteUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMain(m *testing.M) {
    result := m.Run()
    os.Exit(result)
}
