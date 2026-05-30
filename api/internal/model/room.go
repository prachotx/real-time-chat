package model

type Room struct {
	GormModel
	Name        string `gorm:"not null; unique" json:"name"`
	CreatedByID uint   `gorm:"not null" json:"created_by_id"`
	CreatedBy   User   `gorm:"foreignKey:CreatedByID" json:"created_by"`
}
