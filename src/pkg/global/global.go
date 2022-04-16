package global

import (
	"github.com/nasanasaQuQ/goProject/src/conf"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	CONFIG           conf.Config
	VP               *viper.Viper
	LOG              *zap.SugaredLogger
	OFFICIAL_ACCOUNT *officialaccount.OfficialAccount
)
