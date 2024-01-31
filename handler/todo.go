package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dongurikoko/GO_lesson1/usecase"

	"github.com/gin-gonic/gin"
)

type TodoHandler interface {
	HandleTodoCreate(c *gin.Context)
	HandleTodoGet(c *gin.Context)
	HandleTodoUpdate(c *gin.Context)
	HandleTodoDelete(c *gin.Context)
}

type todoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(tu usecase.TodoUseCase) TodoHandler {
	return &todoHandler{
		todoUseCase: tu,
	}
}

// TODOの新規作成
func (th todoHandler) HandleTodoCreate(c *gin.Context) {
	// フォームデータから "title" を取得
	title := c.PostForm("title")

	/* JSONリクエストボディの場合は以下のようになる
	type Todo struct {
		Title string `json:"title"`
	}
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
	*/

	// titleを元にTodoを作成
	if err := th.todoUseCase.Insert(title); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service Unavailable"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})

}

// TODOの取得
func (th todoHandler) HandleTodoGet(c *gin.Context) {
	type todoResponse struct {
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Updated_at time.Time `json:"updated_at"`
	}
	type todosResponse struct {
		Todos []*todoResponse `json:"todos"`
	}

	// クエリパラメータから "query" を取得
	query := c.Query("title")

	// クエリを元にTodoを取得
	todos, err := th.todoUseCase.GetAllByQuery(query)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service is Unavailable"})
		return
	}

	// 取得したドメインモデルをレスポンスモデルに変換
	var response todosResponse
	for _, todo := range todos {
		response.Todos = append(response.Todos, &todoResponse{
			ID:         todo.ID,
			Title:      todo.Title,
			Updated_at: todo.Updated_at,
		})
	}

	//c.JSON(http.StatusOK, response)

	// HTMLテンプレートにデータを渡してレンダリング
	c.HTML(http.StatusOK, "index.html", response)
}

// TODOの更新
func (th todoHandler) HandleTodoUpdate(c *gin.Context) {
	// パスパラメータから "id" を取得
	stringID := c.Param("id")
	// idをint型に変換
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service Unavailable"})
		return
	}

	// フォームデータから "title" を取得
	title := c.PostForm("title")

	// idを元にTodoを更新
	if err := th.todoUseCase.Update(id, title); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 指定されたTODOが存在しなかった場合
			c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		} else {
			// それ以外のエラーの場合
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service Unavailable"})
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
}

// TODOの削除
func (th todoHandler) HandleTodoDelete(c *gin.Context) {
	// パスパラメータから "id" を取得
	stringID := c.Param("id")
	// idをint型に変換
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service Unavailable"})
		return
	}

	// idを元にTodoを削除
	if err := th.todoUseCase.Delete(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 指定されたTODOが存在しなかった場合
			c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		} else {
			// それ以外のエラーの場合
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service Unavailable"})
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "No Content"})
}
