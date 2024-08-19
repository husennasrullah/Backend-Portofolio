package out

import (
	"crudproduct/util"
)

type Payload struct {
	Status StatusResponse `json:"status"`
	Data   interface{}    `json:"data"`
}

type StatusResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (ar Payload) String() string {
	return util.StructToJSON(ar)
}
