package controllers

import (
    "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
    "bikrent/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	// Parse user registration data from the request
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Check if the username is already taken
	var existingUser models.User
	if err := uc.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, "Username already in use")
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := uc.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	// Return a success response
	return c.JSON(http.StatusCreated, "User created successfully")
}

func (uc *UserController) GetUsers(c echo.Context) error {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}

	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) {
    // Logic untuk memperbarui informasi pengguna
}

func DeleteUser(c echo.Context) {
    // Logic untuk menghapus pengguna
}
