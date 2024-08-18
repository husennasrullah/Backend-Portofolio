package queryparam

import (
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/utils"
)

const (
	ParamPage    = "page"
	ParamLimit   = "limit"
	ParamOrderBy = "order_by"
	ParamSort    = "sort"
)

type Param struct {
	Page    int               `json:"page"`
	Limit   int               `json:"limit"`
	Offset  int               `json:"offset"`
	OrderBy string            `json:"order_by"`
	Filter  map[string]string `json:"filter"`
}

func RequestParam(c *gin.Context) (Param, map[string]string) {
	var (
		p            Param
		defaultPage  = 1
		defaultLimit = 10
	)
	queryParams := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			switch key {
			case ParamPage:
				p.Page = utils.ConvertStrToInt(values[0], defaultPage)
			case ParamLimit:
				p.Limit = utils.ConvertStrToInt(values[0], defaultLimit)
			case ParamOrderBy:
				p.OrderBy = values[0]
			default:
				queryParams[key] = values[0]
			}
		}
	}
	if p.Page == 0 {
		p.Page = defaultPage
	}

	if p.Limit == 0 {
		p.Limit = defaultLimit
	}
	p.Filter = queryParams
	p.Offset = (p.Page - 1) * p.Limit

	// Mengambil header
	headers := make(map[string]string)
	for key, values := range c.Request.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}
	return p, headers
}
