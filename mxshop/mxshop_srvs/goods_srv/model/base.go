package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GormList []string

// 将数据转换成json字符串保存在数据库中，当将数据保存到数据库时自动调用
func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 读取数据库中的数据，将json转换成GormList，当从数据库读取数据时自动调用
func (g *GormList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), g)
}

type BaseModel struct {
	ID        int32     `gorm:"primary_key;type:int"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDelete  bool
}
