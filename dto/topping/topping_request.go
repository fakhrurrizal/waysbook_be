package toppingdto

type CreateTopping struct {
	ID    int    `json:"id"`
	Title string `json:"title"  gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type : int" validate:"required"`
	Image string `json:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
}

type UpdateTopping struct {
	ID    int    `json:"id"`
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: int" `
	Image string `json:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
}
