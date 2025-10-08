package service

import (
	"errors"
	"time"
)

type Reward struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	StockID   int       `json:"stock_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time  `json:"timestamp"`
}

type RewardService struct {
	rewards []Reward
}

func NewRewardService() *RewardService {
	return &RewardService{
		rewards: []Reward{},
	}
}

func (s *RewardService) RecordReward(userID int, stockID int, amount float64) (Reward, error) {
	if userID <= 0 || stockID <= 0 || amount <= 0 {
		return Reward{}, errors.New("invalid input parameters")
	}

	reward := Reward{
		ID:        len(s.rewards) + 1,
		UserID:    userID,
		StockID:   stockID,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	s.rewards = append(s.rewards, reward)
	return reward, nil
}

func (s *RewardService) GetRewardsByUser(userID int) ([]Reward, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	var userRewards []Reward
	for _, reward := range s.rewards {
		if reward.UserID == userID {
			userRewards = append(userRewards, reward)
		}
	}

	return userRewards, nil
}

func (s *RewardService) CalculateTotalRewards(userID int) (float64, error) {
	if userID <= 0 {
		return 0, errors.New("invalid user ID")
	}

	total := 0.0
	for _, reward := range s.rewards {
		if reward.UserID == userID {
			total += reward.Amount
		}
	}

	return total, nil
}