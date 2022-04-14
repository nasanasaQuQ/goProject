package global

import (
	"github.com/nasanasaQuQ/goProject/src/conf"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG conf.Config
)
