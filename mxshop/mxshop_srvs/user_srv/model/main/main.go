package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

// 生成md5
func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	//dsn := "root:1234@tcp(127.0.0.1:3307)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//// 设置全局logger，在执行sql语句的时候会打印出来
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,        // Don't include params in the SQL log
	//		Colorful:                  true,        // Disable color
	//	},
	//)
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	//NamingStrategy: schema.NamingStrategy{
	//	//	SingularTable: true,
	//	//	TablePrefix:   "mxshop_",
	//	//	NameReplacer:  nil,
	//	//},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// 建表
	//_ = db.AutoMigrate(&model.User{})
	//fmt.Println(genMd5("123456"))

	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	options := &password.Options{10, 100, 20, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	password1 := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(password1)
	fmt.Println(len(password1))
	passwordInfo := strings.Split(password1, "$")
	fmt.Println(passwordInfo)
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
}
