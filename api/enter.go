package api

import "blogx_server/api/site_api"

type Api struct {
	SiteApi site_api.SiteApi
}

var App = Api{}
