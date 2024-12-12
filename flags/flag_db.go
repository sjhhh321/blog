package flags

import (
	"blogx_server/global"
	"blogx_server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{})
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功!")
}
