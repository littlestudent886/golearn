package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	dsn := "root:1234@tcp(127.0.0.1:3307)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//	TablePrefix:   "mxshop_",
		//	NameReplacer:  nil,
		//},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

}
