package controllers

import (
    "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
    "bikrent/models"
	"log"
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
	inputUsername := c.FormValue("username")
    inputPassword := c.FormValue("password")

    // Check if the username is already taken.
    var existingUser models.User
    if err := uc.DB.Where("username = ?", inputUsername).First(&existingUser).Error; err == nil {
        return c.JSON(http.StatusConflict, map[string]string{"error": "Username already taken"})
    }

    // Hash the password.
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
    }

    // Create a new user.
    newUser := models.User{
        Username: inputUsername,
        Password: string(hashedPassword),
		Role:     0,
    }

    if err := uc.DB.Create(&newUser).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user", "details": err.Error()})
    }    

    return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (uc *UserController) GetUsers(c echo.Context) error {
    var users []models.User
    if err := uc.DB.Find(&users).Error; err != nil {
        log.Println("Error:", err) // Log the actual error
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
