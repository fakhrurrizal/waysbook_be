package transactiondto

type TransactionRequest struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	ToppingID []int  `json:"topping_id"`
	QTY       int    `json:"qty"`
	SubTotal  int    `json:"subtotal"`
	Status    string `jsom:"status"`
}

type TransactionResponse struct {
	ID      int      `json:"id"`
	Name    string   `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email   string   `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone   string   `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address string   `json:"address" form:"address" gorm:"type : varchar(255)"`
	UserID  int      `json:"user_id"`
	Order   []string `json:"order"`
	Status  string   `jsom:"status"`
	Total   int      `json:"total" gorm:"type : int"`
}
