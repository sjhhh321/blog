package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()

	global.Config = core.ReadConf()
	core.InitLogrus()
	logrus.Warnf("xxx")
	logrus.Debug("yyy")
	logrus.Error("zzz")
	logrus.Infof("11123")
}
