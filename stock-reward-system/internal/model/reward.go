package model

type Reward struct {
    ID          int     `json:"id"`
    UserID      int     `json:"user_id"`
    StockSymbol string  `json:"stock_symbol"`
    Amount      float64 `json:"amount"`
    CreatedAt   string  `json:"created_at"`
}

func NewReward(userID int, stockSymbol string, amount float64, createdAt string) *Reward {
    return &Reward{
        UserID:      userID,
        StockSymbol: stockSymbol,
        Amount:      amount,
        CreatedAt:   createdAt,
    }
}