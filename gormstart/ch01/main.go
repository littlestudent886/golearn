package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:1234@tcp(127.0.0.1:3307)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局logger，在执行sql语句的时候会打印出来
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 建表
	//_ = db.AutoMigrate(&Product{})

	// Create 增
	db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})

	// Read 查
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// 改
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	// Update语句会更新零值，Updates语句不会更新零值
	// 1.指针可以更新零值
	// 2.sql.NullString可以更新零值
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 删 执行的实际上是update语句 逻辑删除
	// Delete - delete product
	//db.Delete(&product, 1)
}
