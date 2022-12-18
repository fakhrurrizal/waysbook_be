package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}