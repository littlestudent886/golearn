package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/user_srv/config"
)

var DB *gorm.DB
var ServerConfig *config.ServerConfig = &config.ServerConfig{}
