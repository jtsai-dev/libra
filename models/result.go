/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:16:49
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 14:24:43
 */
package models

import (
	"time"

	"libra/pkg/enums"
)

type Result struct {
	Success    bool        `json:"success"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	ServerTime time.Time   `json:"serverTime"`
}

func R(data interface{}) *Result {
	model := Result{
		Success:    true,
		Code:       enums.Success,
		Message:    enums.GetRespCodeDesc(enums.Success),
		Data:       &data,
		ServerTime: time.Now(),
	}
	model.Message = enums.GetRespCodeDesc(model.Code)
	return &model
}

func RC(code int) *Result {
	model := Result{
		Success:    code == enums.Success,
		Code:       code,
		Message:    enums.GetRespCodeDesc(code),
		ServerTime: time.Now(),
	}
	return &model
}

func RCM(code int, message string) *Result {
	model := Result{
		Success:    code == enums.Success,
		Code:       code,
		Message:    message,
		ServerTime: time.Now(),
	}

	if len(message) == 0 {
		model.Message = enums.GetRespCodeDesc(code)
	}

	return &model
}
