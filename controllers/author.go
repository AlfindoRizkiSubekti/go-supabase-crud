package controllers

import (
	"context"
	"net/http"

	"go-supabase-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthorController struct {
	DB *pgxpool.Pool
}

func NewAutorController(db *pgxpool.Pool) *AuthorController {
	return &AuthorController{DB: db}
}

func (bc *AuthorController) GetAuthors(c *gin.Context) {
	// Gunakan context khusus
	ctx := context.WithValue(c.Request.Context(), "pgx.DefaultQueryExecMode", "simple")

	query := `SELECT authors.name
              FROM authors`

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

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(author.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Data parsing error",
				"details": err.Error(),
			})
			return
		}
		authors = append(authors, author)
	}

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}
