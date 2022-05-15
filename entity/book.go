package entity

// Book represents a book in database .
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserID      uint64 `gorm:"type:int" json:"user_id"`
	User        User   `gorm:"foreignkey:UserID;contraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
