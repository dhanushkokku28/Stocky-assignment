package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stock-reward-system/internal/model"
	"stock-reward-system/internal/service"
)

// RewardHandler handles reward-related API requests
type RewardHandler struct {
	rewardService service.RewardService
}

// NewRewardHandler creates a new RewardHandler
func NewRewardHandler(rewardService service.RewardService) *RewardHandler {
	return &RewardHandler{rewardService: rewardService}
}

// RecordReward records a new reward
func (h *RewardHandler) RecordReward(c *gin.Context) {
	var reward model.Reward
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.rewardService.RecordReward(reward); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record reward"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reward recorded successfully"})
}

// GetTodaysStocks retrieves today's stock values
func (h *RewardHandler) GetTodaysStocks(c *gin.Context) {
	stocks, err := h.rewardService.GetTodaysStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch today's stocks"})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// GetHistoricalINR retrieves historical INR values
func (h *RewardHandler) GetHistoricalINR(c *gin.Context) {
	historicalData, err := h.rewardService.GetHistoricalINR()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch historical INR values"})
		return
	}

	c.JSON(http.StatusOK, historicalData)
}

// GetUserStats retrieves user statistics
func (h *RewardHandler) GetUserStats(c *gin.Context) {
	userID := c.Param("user_id")
	stats, err := h.rewardService.GetUserStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}