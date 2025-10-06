package category

import "time"

type Category struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;not null;unique"`
	Description string    `gorm:"column:description;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
