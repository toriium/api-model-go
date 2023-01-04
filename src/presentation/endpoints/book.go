package endpoints

import (
	bookService "api/src/services/book"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ConsultBookInput struct {
	BookId string `form:"book_id" binding:"required"`
}

type ConsultBookOutput struct {
	BookId      int       `json:"book_id" binding:"required"`
	Isbn        string    `json:"isbn" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Publisher   string    `json:"publisher" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required" time_format:"01/02/06"`
	Pages       int       `json:"pages" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

// ConsultBook insere um usuário no banco de dados
func ConsultBook(c *gin.Context) {
	var validation ConsultBookInput
	if err := c.ShouldBindQuery(&validation); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	service := bookService.BookService{}

	bookValue, _ := service.FindBookById(45)

	response := ConsultBookOutput{
		BookId:      bookValue.BookId,
		Isbn:        bookValue.Isbn,
		Name:        bookValue.Name,
		Author:      bookValue.Author,
		Publisher:   bookValue.Publisher,
		ReleaseDate: bookValue.ReleaseDate,
		Pages:       bookValue.Pages,
		Description: bookValue.Description,
	}

	c.SecureJSON(http.StatusOK, response)
}
