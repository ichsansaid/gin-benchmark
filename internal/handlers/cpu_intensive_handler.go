package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ichsansaid/golang-gin-benchmark/internal/contracts"
)

// CPUIntensiveHandler handles CPU-intensive tasks
type CPUIntensiveHandler struct {
	fiboDpCpuIntensiveUcase        contracts.CPUntensiveUcaseContract
	fiboRecursiveCpuIntensiveUcase contracts.CPUntensiveUcaseContract
}

// NewCPUIntensiveHandler creates a new instance of CPUIntensiveHandler
func NewCPUIntensiveHandler(fiboDpCpuIntensiveUcase contracts.CPUntensiveUcaseContract, fiboRecursiveCpuIntensiveUcase contracts.CPUntensiveUcaseContract) *CPUIntensiveHandler {
	return &CPUIntensiveHandler{
		fiboDpCpuIntensiveUcase:        fiboDpCpuIntensiveUcase,
		fiboRecursiveCpuIntensiveUcase: fiboRecursiveCpuIntensiveUcase,
	}
}

// HandleFibonaciDP handles a CPU-intensive Fibonacci Dynamic Programming task request
func (h *CPUIntensiveHandler) HandleFibonaciDP(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := h.fiboDpCpuIntensiveUcase.DoTask(context.Background(), n)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

// HandleFibonaciRecursive handles a CPU-intensive Fibonacci Recursive task request
func (h *CPUIntensiveHandler) HandleFibonaciRecursive(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := h.fiboRecursiveCpuIntensiveUcase.DoTask(context.Background(), n)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
