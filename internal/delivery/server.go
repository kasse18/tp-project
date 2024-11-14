package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"tp-project/internal/delivery/handlers"
	"tp-project/internal/delivery/handlers/middleware"
	repository "tp-project/internal/repository"
	service "tp-project/internal/service"
	"tp-project/pkg/logger"
)

func Start(db *sqlx.DB, logger *logger.Logger) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	mdw := middleware.InitMiddleware(logger)
	r.Use(mdw.CORSMiddleware())
	r.Use(mdw.AuthMiddleware())

	userRepo := repository.InitUserRepo(db, *logger)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRouter := r.Group("/user")

	userRouter.GET("/:id", mdw.CORSMiddleware(), mdw.AuthMiddleware(), userHandler.GetUserByID)
	userRouter.GET("/email/:email", mdw.CORSMiddleware(), mdw.AuthMiddleware(), userHandler.GetUserByEmail)
	userRouter.POST("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), userHandler.CreateUser)
	userRouter.PUT("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), userHandler.UpdateUser)
	userRouter.DELETE("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), userHandler.DeleteUser)

	tattooRepo := repository.InitTattooRepo(db, *logger)
	tattooService := service.NewTattooService(tattooRepo)
	tattooHandler := handlers.NewTattooHandler(tattooService)

	tattooRouter := r.Group("/tattoo")

	tattooRouter.GET("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), tattooHandler.GetAllTattoos)
	tattooRouter.GET("/:id", mdw.CORSMiddleware(), mdw.AuthMiddleware(), tattooHandler.GetTattooByID)
	tattooRouter.POST("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), tattooHandler.CreateTattoo)
	tattooRouter.PUT("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), tattooHandler.UpdateTattoo)
	tattooRouter.DELETE("/", mdw.CORSMiddleware(), mdw.AuthMiddleware(), tattooHandler.DeleteTattoo)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
