package controllers

import (
    "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
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

    var existingUser models.User
    if err := uc.DB.Where("username = ?", inputUsername).First(&existingUser).Error; err == nil {
        return c.JSON(http.StatusConflict, map[string]string{"error": "Username already taken"})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
    }

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

func (uc *UserController) Login(c echo.Context) error {
    inputUsername := c.FormValue("username")
    inputPassword := c.FormValue("password")

    var user models.User
    if err := uc.DB.Where("username = ?", inputUsername).First(&user).Error; err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword)); err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
    }

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = user.Username
    claims["userID"] = user.ID
    claims["role"] = user.Role
    tokenString, err := token.SignedString([]byte("your-secret-key")) 

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT token"})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func (uc *UserController) GetUsers(c echo.Context) error {
    var users []models.User
    if err := uc.DB.Find(&users).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
    }

    return c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c echo.Context) error {
    userID := c.Param("id")

    var user models.User
    if err := uc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    // Fetch rental history for the user and include it in the response
    var rentals []models.Rental
    if err := uc.DB.Where("user_id = ?", user.ID).Find(&rentals).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user rentals"})
    }
    user.Rentals = rentals

    return c.JSON(http.StatusOK, user)
}


func (uc *UserController) UpdateUser(c echo.Context) error {
    userID := c.Param("id")

    var user models.User
    if err := uc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    updatedUser := new(models.User)
    if err := c.Bind(updatedUser); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if updatedUser.Username != "" {
        user.Username = updatedUser.Username
    }
    if updatedUser.Password != "" {
        user.Password = updatedUser.Password
    }

    if err := uc.DB.Save(&user).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
    }

    return c.JSON(http.StatusOK, user)
}


func (uc *UserController) DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	if err := uc.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}