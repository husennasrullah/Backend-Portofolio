package dto

import (
	"crudproduct/model"
	"strings"
)

type GetlistDTO struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"order_by"`
}

func (input *GetlistDTO) ValidatePageLimitAndOrderBy() (result GetlistDTO, err model.ErrorModel) {
	//funcName := "ValidatePageLimitAndOrderBy"
	if input.Page < 0 {
		err = model.GenerateErrorModel(400, "PAGE_NEED_MORE_THAN", "pagination.go", "ValidatePageLimitAndOrderBy", nil)
		return
	}

	if input.Limit < 0 {
		err = model.GenerateErrorModel(400, "LIMIT_NEED_MORE_THAN", "pagination.go", "ValidatePageLimitAndOrderBy", nil)
		return
	}

	var isAscending = true

	input.OrderBy = strings.Trim(input.OrderBy, " ")
	if input.OrderBy != "" {
		tempOrder := strings.Split(input.OrderBy, " ")
		input.OrderBy = tempOrder[0]

		if len(tempOrder) > 1 {
			if strings.ToLower(tempOrder[1]) == "desc" {
				isAscending = false
			}
		}

	} else {
		//set default order by
		input.OrderBy = "id"
	}

	if isAscending {
		input.OrderBy += " ASC"
	} else {
		input.OrderBy += " DESC"
	}

	result = *input
	return
}
