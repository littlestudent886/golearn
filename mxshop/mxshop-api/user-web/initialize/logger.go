package initialize

import "go.uber.org/zap"

/*
初始化日志 替换全局logger
1.S()可以获取一个全局的sugar，可以自己设置一个全局的logger
2.日志是分级别的，debug，info，warn，error，panic，fatal
3.S函数和L函数很有用，可以提供一个全局的安全访问logger的途径
*/
func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	zap.ReplaceGlobals(logger)
}
