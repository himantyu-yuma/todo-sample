package main

import (
	"log"

	"todo-sample/config"
	"todo-sample/infrastructure/persistence"
	"todo-sample/infrastructure/persistence/model"

	"todo-sample/interface/handler"
	"todo-sample/usecase"

	"todo-sample/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()

	db, err := persistence.NewDB(*config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	todoRepo := model.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)
	healthCheckHandler := handler.NewHealthCheckHandler(db)

	r := gin.Default()

	r.GET("/health", healthCheckHandler.HealthCheck)

	r.POST("/todos", todoHandler.CreateTodo)

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Description = "This is a sample server Todo server."
	docs.SwaggerInfo.Version = "1.0"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
