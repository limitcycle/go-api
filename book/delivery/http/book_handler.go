package http

import (
	"context"
	"fmt"
	bookUcase "go-api/book"
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpBookHandler struct {
	BUsecase bookUcase.BookUsecase
}

func (b *HttpBookHandler) Fetch(c *gin.Context, cursor string, num int) {
	// numP, err := strconv.Atoi(c.Query("num"))
	// num := int64(numP)
	// cursor := c.Query("cursor")
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listBook, nextCursor, err := b.BUsecase.Fetch(ctx, cursor, num)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.Header(`X-Cursor`, nextCursor)
	c.JSON(http.StatusOK, listBook)
}

func (b *HttpBookHandler) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	book, err := b.BUsecase.GetByID(ctx, idP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
func (b *HttpBookHandler) GetRouter(c *gin.Context) {
	author := c.Query("author")
	name := c.Query("name")
	idP, _ := strconv.Atoi(c.Query("num"))

	cursor := c.Query("cursor")
	fmt.Printf("author = %s, name = %s, id = %d, cursor = %s", author, name, idP, cursor)

	if cursor != "" && idP != 0 {
		b.Fetch(c, cursor, idP)
	} else if author != "" {
		b.GetByAuthor(c, author)
	} else if name != "" {
		b.GetByName(c, name)
	} else {
		c.JSON(http.StatusBadRequest, models.INTERNAL_SERVER_ERROR)
	}
}
func (b *HttpBookHandler) GetByAuthor(c *gin.Context, author string) {
	// authorName := c.Param("name")
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	book, err := b.BUsecase.GetByAuthor(ctx, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *HttpBookHandler) GetByName(c *gin.Context, name string) {
	// bookName := c.Param("name")
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	book, err := b.BUsecase.GetByName(ctx, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *HttpBookHandler) Store(c *gin.Context) {
	var book models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	saveBook, err := b.BUsecase.Store(ctx, &book)

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, saveBook)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (b *HttpBookHandler) Update(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))

	var book models.Book
	book.ID = idP
	err = c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		return
	}
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	updateBook, err := b.BUsecase.Update(ctx, &book)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, updateBook)
}

func (b *HttpBookHandler) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	_, err = b.BUsecase.Delete(ctx, idP)

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func NewBookHttpHandler(r *gin.Engine, us bookUcase.BookUsecase) {
	handler := &HttpBookHandler{
		BUsecase: us,
	}
	v1 := r.Group("api/v1")
	v1.GET("/book", handler.GetRouter)
	v1.GET("/book/:id", handler.GetByID)
	//	v1.GET("/book/author", handler.GetByAuthor)
	//	v1.GET("/book/name", handlers ...gin.HandlerFunc)
	v1.POST("/book", handler.Store)
	v1.PUT("/book/:id", handler.Update)
	v1.DELETE("/book/:id", handler.Delete)
}
