package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichsansaid/golang-gin-benchmark/internal/contracts"
)

// IOIntensiveHandler handles IO-intensive tasks
type IOIntensiveHandler struct {
	googleIOIntensiveUcase contracts.IOIntensiveUcaseContract
}

// NewIOIntensiveHandler creates a new instance of IOIntensiveHandler
func NewIOIntensiveHandler(googleIOIntensiveUcase contracts.IOIntensiveUcaseContract) *IOIntensiveHandler {
	return &IOIntensiveHandler{
		googleIOIntensiveUcase: googleIOIntensiveUcase,
	}
}

// HandleGoogleIO handles an IO-intensive task request
func (h *IOIntensiveHandler) HandleGoogleIO(c *gin.Context) {
	h.googleIOIntensiveUcase.DoTask(context.Background())
	c.JSON(http.StatusOK, gin.H{"message": "Google IO task completed"})
}
