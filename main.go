package main

import (
	"diitfin/config"
	"diitfin/controller"
	"diitfin/model"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	if err := gorm.AutoMigrate(&model.Users{}); err != nil {
		panic(err)
	}

	dbGorm.Ping()

	userRoute := e.Group("/user")
	userRoute.GET("/all", controller.GetAllUser)
	userRoute.POST("/create", controller.AddUser)
	userRoute.GET("/:id", controller.UserByID)
	userRoute.PUT("/update/:id", controller.UpdateUser)
	userRoute.DELETE("/delete/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
