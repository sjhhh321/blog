package site_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}
