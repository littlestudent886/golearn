package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	//db.AutoMigrate(&User1{})

	//languages := []Language{{Name: "English"}, {Name: "Chinese"}}
	//user := User1{Languages: languages}
	//db.Create(&user)

	//var user User1
	//db.Preload("Languages").First(&user)
	//for _, language := range user.Languages {
	//	println(language.Name)
	//}

	// 如果之前已经取出来一个用户，但是之前没有用preload来加载对应的语言，那么可以通过Association来加载
	var user User1
	db.First(&user)
	var languages []Language
	db.Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		println(language.Name)
	}

}
