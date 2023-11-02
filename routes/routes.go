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
    // authenticated := e.Group("/api", middleware.InitJWTMiddleware(db))
    
    //users
    e.POST("/users/register", userController.CreateUser)
    e.POST("/users/login", userController.Login)
    e.GET("/users", middleware.AdminMiddleware(userController.GetUsers))
    e.GET("/users/:id", userController.GetUserByID)
    e.PUT("/users/:id", userController.UpdateUser)
    e.DELETE("/users/:id", middleware.SuperAdminMiddleware(userController.DeleteUser))

    //bicycle
    e.GET("/bicycles", bikeController.GetBicycles)
    e.GET("/bicycles/:id", bikeController.GetBicycleByID)
    e.POST("/bicycles/register", middleware.AdminMiddleware(bikeController.CreateBicycle))
    e.PUT("/bicycles/:id", middleware.AdminMiddleware(bikeController.UpdateBicycle))
    e.DELETE("/bicycles/:id", middleware.SuperAdminMiddleware(bikeController.DeleteBicycle))

    //userdetail
    e.POST("/userdetail/:id", userDetailController.CreateUserDetail)
    e.GET("/userdetail/:id", userDetailController.GetUserWithDetail)
    e.PUT("userdetail/:id", userDetailController.UpdateUserDetail)
    e.DELETE("userdetail/:id", middleware.SuperAdminMiddleware(userDetailController.DeleteUserDetail))

    //rental
    e.GET("/rental/:id", rentalController.GetRental)
	e.POST("/rentals", rentalController.CreateRental)

    e.POST("/getbikefact",getbikefact)
}
