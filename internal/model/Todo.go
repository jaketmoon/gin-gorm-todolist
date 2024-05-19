package model

import (
	"gorm.io/gorm"
	"time"
)

// 注意：所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
//如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口

type Todo struct {
	ID        uint           `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `gorm:"varchar(20);not null" json:"title"` //标题
	Status    bool           `json:"status"`                            //状态
	//Userid string `gorm:"varchar(20);not null" json:"userid"` //用户id
}
