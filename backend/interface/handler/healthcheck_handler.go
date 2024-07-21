package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheckHandler struct {
	DB *gorm.DB
}

func NewHealthCheckHandler(db *gorm.DB) *HealthCheckHandler {
	return &HealthCheckHandler{
		DB: db,
	}
}

func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	sqlDB, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection error"})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database ping error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server is healthy"})
}
