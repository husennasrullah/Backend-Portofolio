package model

import "errors"

type ErrorModel struct {
	Code     int
	Error    error
	FileName string
	FuncName string
	CausedBy error
}

func GenerateErrorModel(code int, err string, fileName string, funcName string, causedBy error) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	errModel.CausedBy = causedBy
	return errModel
}
