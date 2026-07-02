package models

type GirlModel struct {
	ID      int        `gorm:"primary_key"`
	Name    string     `gorm:"size:20"`
	BoyList []BoyModel `gorm:"foreignKey:GirlID"`
}
type BoyModel struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"size:20"`
	GirlID    int
	GirlModel *GirlModel `gorm:"foreignKey:GirlID"`
}
