package controllers

import (
    "github.com/labstack/echo"
	// "golang.org/x/crypto/bcrypt"
    // "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
    "bikrent/models"
	"log"
)

type UserDetailController struct {
	DB *gorm.DB
}

func NewUserDetailController(db *gorm.DB) *UserDetailController {
	return &UserDetailController{
		DB: db,
	}
}

func (udc *UserDetailController) CreateUserDetail(c echo.Context) error {
    userID := c.Param("id")

    var user models.User
    if err := udc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    nama := c.FormValue("nama")
    alamat := c.FormValue("alamat")
    noTelp := c.FormValue("notelp")
    jenisKelamin := c.FormValue("jeniskelamin")
    tanggalTempatLahir := c.FormValue("tempattanggallahir")

    userDetail := models.UserDetail{
        UserID:             user.ID,
        Nama:               nama,
        Alamat:             alamat,
        NoTelp:             noTelp,
        JenisKelamin:       jenisKelamin,
        TanggalTempatLahir: tanggalTempatLahir,
    }

    if err := udc.DB.Create(&userDetail).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user detail"})
    }

    user.UserDetail = userDetail
    if err := udc.DB.Save(&user).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user with user detail"})
    }

    return c.JSON(http.StatusCreated, userDetail)
}

func (udc *UserDetailController) GetUserWithDetail(c echo.Context) error {
    userID := c.Param("id")

    var user models.User

    if err := udc.DB.Preload("UserDetail").First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    return c.JSON(http.StatusOK, user)
}

func (udc *UserDetailController) UpdateUserDetail(c echo.Context) error {
    userID := c.Param("id")

    // Find the user by ID
    var user models.User
    if err := udc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    // Find the user detail by user ID
    var userDetail models.UserDetail
    if err := udc.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "UserDetail not found"})
    }

    // Parse the fields you want to update from the request
    nama := c.FormValue("nama")
    alamat := c.FormValue("alamat")
    noTelp := c.FormValue("notelp")
	jeniskelamin := c.FormValue("jeniskelamin")
	tempattanggallahir := c.FormValue("tempattanggallahir")

    // Check if each field is updated and assign the updated value or keep the existing data
    if nama != "" {
        userDetail.Nama = nama
    }
    if alamat != "" {
        userDetail.Alamat = alamat
    }
    if noTelp != "" {
        userDetail.NoTelp = noTelp
    }
	if jeniskelamin != "" {
        userDetail.NoTelp = jeniskelamin
    }
	if tempattanggallahir != "" {
        userDetail.NoTelp = tempattanggallahir
    }

    if err := udc.DB.Save(&userDetail).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update UserDetail"})
    }

    return c.JSON(http.StatusOK, userDetail)
}


func (udc *UserDetailController) DeleteUserDetail(c echo.Context) error {
    userID := c.Param("id")

    var user models.User
    if err := udc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    var userDetail models.UserDetail
    if err := udc.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "UserDetail not found"})
    }

    if err := udc.DB.Delete(&userDetail).Error; err != nil {
        log.Println("Error:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete UserDetail"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "UserDetail deleted successfully"})
}

func (udc *UserDetailController) GetUserDetailByID(c echo.Context) error {
    // Get the user ID from the URL parameter
    userID := c.Param("id")

    // Find the user by ID
    var user models.User
    if err := udc.DB.First(&user, userID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    // Find the user detail by user ID
    var userDetail models.UserDetail
    if err := udc.DB.Where("user_id = ?", user.ID).First(&userDetail).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "UserDetail not found"})
    }

    return c.JSON(http.StatusOK, userDetail)
}