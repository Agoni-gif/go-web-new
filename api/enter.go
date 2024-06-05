package api

import v1 "go-web-new/api/v1"

type ApiGroup struct {
	V1Group v1.V1Group
}

var ApiGroupApp = new(ApiGroup)
