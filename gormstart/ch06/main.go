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

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         `gorm:"column:myname"` // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
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

	var user User
	//// Get first matched record
	//db.Where("name = ?", "jinzhu").First(&user)                     //如果变量名和数据库列名不一致，则需要使用列名
	//db.Where(&User{Name: "jinzhu"}).First(&user)                    // 或者使用结构体
	//db.Where(map[string]interface{}{"Name": "jinzhu"}).First(&user) //还有map用法 map和第一种方法一样，得用列名
	////db.Where("name = ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	//
	//var users []User
	//// Get all matched records
	//db.Where("name <> ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name <> 'jinzhu';
	//
	//db.Where([]int64{20, 21, 22}).Find(&users) //主键切片条件

	db.Where("age = ?", 0).First(&user)
	db.Where(&User{Name: "jinzhu", Age: 0}).First(&user)                      // 结构体方式不能查询0
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).First(&user) // map方式可以查询0

}
