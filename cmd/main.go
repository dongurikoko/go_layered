package main

import (
	"log"

	"github.com/dongurikoko/GO_lesson1/config"
	"github.com/dongurikoko/GO_lesson1/handler"
	"github.com/dongurikoko/GO_lesson1/infra/persistence"
	"github.com/dongurikoko/GO_lesson1/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// 依存関係を注入
	db, _ := config.GetConn()
	todoPersistence := persistence.NewTodoPersistence(db)
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodoHandler(todoUseCase)

	engine := gin.Default()
	// htmlのディレクトリを指定
	engine.LoadHTMLGlob("public/*.html")

	engine.POST("/todo/create", todoHandler.HandleTodoCreate)
	engine.GET("/todo/get", todoHandler.HandleTodoGet)
	engine.POST("/todo/update/:id", todoHandler.HandleTodoUpdate)
	engine.DELETE("/todo/delete/:id", todoHandler.HandleTodoDelete)

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	engine.Run(":8080")
}
