package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/goods_srv/config"
)

var DB *gorm.DB
var ServerConfig *config.ServerConfig = &config.ServerConfig{}
var NacosConfig *config.NacosConfig = &config.NacosConfig{}
