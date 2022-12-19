package productdto

type CreateProduct struct {
	Title string `json:"title" form:"title" gorm:"type : varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type : int" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title" gorm:"type : varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type : int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty   int    `json:"qty" form:"qty"`
}
