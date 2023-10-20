package routes

import (
    "github.com/labstack/echo"
    "bikrent/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userController := controllers.NewUserController(db)
    // Rute untuk entitas pengguna (User)
    // e.POST("/users", controllers.CreateUser)
    e.GET("/users/:id", userController.GetUsers)
    // e.PUT("/users/:id", controllers.UpdateUser)
    // e.DELETE("/users/:id", controllers.DeleteUser)

    // Rute untuk entitas lainnya (Bicycle, Rental, UserDetail)
}
