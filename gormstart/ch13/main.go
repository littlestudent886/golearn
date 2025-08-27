package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User1 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

// 在gorm中可以通过给某一个struct添加tableName方法自定义表名
//func (Language) TableName() string {
//	return "my_language"
//}

func (l *Language) BeforeCreate(tx *gorm.DB) error {
	l.Name = "中文"
	return nil
}

/*
1.自己定义表名
2.统一给表名加前缀
*/
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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop_",
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Language{})

	db.Create(&Language{Name: "中文"})

}
