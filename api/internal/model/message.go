package model

type Message struct {
	GormModel
	Content  string `gorm:"not null" json:"content"`
	RoomID   uint   `gorm:"not null" json:"room_id"`
	Room     Room   `gorm:"foreignKey:RoomID" json:"room"`
	SenderID uint   `gorm:"not null" json:"sender_id"`
	Sender   User   `gorm:"foreignKey:SenderID" json:"sender"`
}
