package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
    "bikrent/models"
	"strconv"
)
func TestUserDetailController_GetUserWithDetail(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    udc := NewUserDetailController(db)

    // Ensure a clean database by deleting all users (if necessary)
    db.Exec("DELETE FROM users")
    db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

    // Define a test user
    testUser := models.User{
        Username: "testuser",
        Password: "hashedpassword",
        Role:     0,
    }
    db.Create(&testUser)

    // Define a test user detail
    testUserDetail := models.UserDetail{
        UserID:             testUser.ID,
        Nama:               "John",
        Alamat:             "123 Main St",
        NoTelp:             "123456789",
        JenisKelamin:       "Male",
        TanggalTempatLahir: "1990-01-01",
    }
    db.Create(&testUserDetail)

    // Define a test request to get user with detail
    req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.FormatUint(uint64(testUser.ID), 10)+"/userdetails", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := udc.GetUserWithDetail(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    // You can further assert the response body to check the returned user with detail.
}

func TestUserDetailController_UpdateUserDetail(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    udc := NewUserDetailController(db)

    // Ensure a clean database by deleting all users (if necessary)
    db.Exec("DELETE FROM users")
    db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

    // Define a test user
    testUser := models.User{
        Username: "testuser",
        Password: "hashedpassword",
        Role:     0,
    }
    db.Create(&testUser)

    // Define a test user detail
    testUserDetail := models.UserDetail{
        UserID:             testUser.ID,
        Nama:               "John",
        Alamat:             "123 Main St",
        NoTelp:             "123456789",
        JenisKelamin:       "Male",
        TanggalTempatLahir: "1990-01-01",
    }
    db.Create(&testUserDetail)

    // Define a test request to update user detail
    formData := "alamat=456 Elm St&notelp=987654321"
    req := httptest.NewRequest(http.MethodPut, "/users/"+strconv.FormatUint(uint64(testUser.ID), 10)+"/userdetails", strings.NewReader(formData))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := udc.UpdateUserDetail(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    // Check if the user detail has been updated in the database
    var updatedUserDetail models.UserDetail
    if err := db.Where("user_id = ?", testUser.ID).First(&updatedUserDetail).Error; err != nil {
        t.Errorf("Expected user detail to be updated for user with ID %d, but got an error: %v", testUser.ID, err)
    }
    assert.Equal(t, "456 Elm St", updatedUserDetail.Alamat)
    assert.Equal(t, "987654321", updatedUserDetail.NoTelp)
    // You can further assert other updated fields.
}

func TestUserDetailController_DeleteUserDetail(t *testing.T) {
    // Create a test Echo instance and a test DB
    e := echo.New()
    db := setupTestDB()
    udc := NewUserDetailController(db)

    // Ensure a clean database by deleting all users (if necessary)
    db.Exec("DELETE FROM users")
    db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

    // Define a test user
    testUser := models.User{
        Username: "testuser",
        Password: "hashedpassword",
        Role:     0,
    }
    db.Create(&testUser)

    // Define a test user detail
    testUserDetail := models.UserDetail{
        UserID:             testUser.ID,
        Nama:               "John",
        Alamat:             "123 Main St",
        NoTelp:             "123456789",
        JenisKelamin:       "Male",
        TanggalTempatLahir: "1990-01-01",
    }
    db.Create(&testUserDetail)

    // Define a test request to delete user detail
    req := httptest.NewRequest(http.MethodDelete, "/users/"+strconv.FormatUint(uint64(testUser.ID), 10)+"/userdetails", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := udc.DeleteUserDetail(c)

    // Assert the response
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    // Check if the user detail has been deleted from the database
    var deletedUserDetail models.UserDetail
    if err := db.Where("user_id = ?", testUser.ID).First(&deletedUserDetail).Error; err == nil {
        t.Errorf("Expected user detail to be deleted for user with ID %d, but it still exists", testUser.ID)
    }
}
