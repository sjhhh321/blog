package core

import (
	ipUtils "blogx_server/utils/ip"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		logrus.Fatalf("ip数据库加载失败: %s", err)
		return
	}
	searcher = _searcher
}
func GetIpAddr(ip string) (addr string) {
	if ipUtils.HasLocalIPAddr(ip) {
		return "内网"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("错误的ip地址 %s", err)
		return "错误的ip地址"
	}
	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		logrus.Warnf("异常的ip地址 %s", ip)
		return "未知地址"
	}
	// addrList ：国家 0  省   市  运营商
	country := _addrList[0]
	province := _addrList[2]
	city := _addrList[3]
	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return country
	}
	return region
}
