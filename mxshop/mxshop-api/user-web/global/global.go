package global

import (
	ut "github.com/go-playground/universal-translator"
	"user-web/config"
)

var (
	//service配置变量
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	//翻译器选项
	Trans ut.Translator
)
