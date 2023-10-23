package routes

import (
    "github.com/labstack/echo"
    "bikrent/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userController := controllers.NewUserController(db)
    bikeController := controllers.NewBicycleController(db)
    userDetailController := controllers.NewUserDetailController(db)
    
    //users
    e.POST("/users/register", userController.CreateUser)
    e.POST("/users/login", userController.Login)
    e.GET("/users", userController.GetUsers)
    e.GET("/users/:id", userController.GetUserByID)
    e.PUT("/users/:id", userController.UpdateUser)
    e.DELETE("/users/:id", userController.DeleteUser)

    //bicycle
    e.GET("/bicycles", bikeController.GetBicycles)
    e.GET("/bicycles/:id", bikeController.GetBicycleByID)
    e.POST("/bicycles/register", bikeController.CreateBicycle)
    e.PUT("/bicycles/:id", bikeController.UpdateBicycle)
    e.DELETE("/bicycles/:id", bikeController.DeleteBicycle)

    //userdetail
    e.POST("/userdetail/:id", userDetailController.CreateUserDetail)
    e.GET("/userdetail/:id", userDetailController.GetUserWithDetail)
    e.PUT("userdetail/:id", userDetailController.UpdateUserDetail)
    e.DELETE("userdetail/:id", userDetailController.DeleteUserDetail)
}
