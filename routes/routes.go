package routes

import (
    "github.com/labstack/echo"
    "bikrent/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userController := controllers.NewUserController(db)
    bikeController := controllers.NewBicycleController(db)
    
    //users
    e.POST("/users/register", userController.CreateUser)
    e.GET("/users", userController.GetUsers)
    // e.PUT("/users/:id", controllers.UpdateUser)
    // e.DELETE("/users/:id", controllers.DeleteUser)

    //bicycle
    e.GET("/bicycles", bikeController.GetBicycles)
    e.POST("/bicycles/register", bikeController.CreateBicycle)
    e.PUT("/bicycles/:id", bikeController.UpdateBicycle)
    e.DELETE("/bicycles/:id", bikeController.DeleteBicycle)
}
