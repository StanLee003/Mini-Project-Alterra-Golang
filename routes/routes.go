package routes

import (
    "github.com/labstack/echo"
    "bikrent/controllers"
    "bikrent/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userController := controllers.NewUserController(db)
    bikeController := controllers.NewBicycleController(db)
    userDetailController := controllers.NewUserDetailController(db)
    rentalController := controllers.NewRentalController(db)
    authenticated := e.Group("/api", middleware.InitJWTMiddleware(db))
    
    //users
    e.POST("/users/register", userController.CreateUser)
    e.POST("/users/login", userController.Login)
    authenticated.GET("/users", userController.GetUsers)
    authenticated.GET("/users/:id", userController.GetUserByID)
    authenticated.PUT("/users/:id", userController.UpdateUser)
    e.DELETE("/users/:id", userController.DeleteUser)

    //bicycle
    e.GET("/bicycles", bikeController.GetBicycles)
    e.GET("/bicycles/:id", bikeController.GetBicycleByID)
    authenticated.POST("/bicycles/register", bikeController.CreateBicycle)
    authenticated.PUT("/bicycles/:id", bikeController.UpdateBicycle)
    authenticated.DELETE("/bicycles/:id", bikeController.DeleteBicycle)

    //userdetail
    authenticated.POST("/userdetail/:id", userDetailController.CreateUserDetail)
    authenticated.GET("/userdetail/:id", userDetailController.GetUserWithDetail)
    authenticated.PUT("userdetail/:id", userDetailController.UpdateUserDetail)
    authenticated.DELETE("userdetail/:id", userDetailController.DeleteUserDetail)

    //rental
    authenticated.GET("/rental/:id", rentalController.GetRental)
	authenticated.POST("/rentals", rentalController.CreateRental)
}
