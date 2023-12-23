package dal

type User struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Gender string `gorm:"column:gender" json:"gender"`
	Phone  string `gorm:"column:phone" json:"phone"`
	Email  string `gorm:"column:email" json:"email"`
}
