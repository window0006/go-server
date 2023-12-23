package dal

type Family struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Patriarch int    `gorm:"column:patriarch" json:"patriarch"`
	Members   string `gorm:"column:members" json:"members"`
}
