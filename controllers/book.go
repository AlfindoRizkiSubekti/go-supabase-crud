package controllers

import (
	"context"
	"net/http"

	"go-supabase-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookController struct {
	DB *pgxpool.Pool
}

func NewBookController(db *pgxpool.Pool) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) GetBooks(c *gin.Context) {
	// Gunakan context khusus
	ctx := context.WithValue(c.Request.Context(), "pgx.DefaultQueryExecMode", "simple")

	query := `SELECT b.id, b.title, b.year, 
                     a.id as author_id, a.name as author_name
              FROM books b
              JOIN authors a ON b.author_id = a.id`

	// Gunakan QueryRow untuk single row atau Query untuk multiple rows
	rows, err := bc.DB.Query(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"details": err.Error(),
		})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Year,
			&book.Author.ID, &book.Author.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Data parsing error",
				"details": err.Error(),
			})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
