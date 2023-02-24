package auth

import (
	"efishery-golang/dto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) Controller {
	return Controller{
		service: service,
	}
}

func (ctrl Controller) Login(body LoginRequest) (res dto.BaseResponse, err error) {
	data, err := ctrl.service.Login(body)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage("Error"), nil
	}
	if data == "" {
		return dto.DefaultErrorBaseResponseWithMessage("Invalid phone or password"), nil
	}
	res = dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Success",
		Data: LoginResponse{
			Token: data,
		},
	}
	return res, nil
}

func (ctrl Controller) register(body RegisterRequest) (res dto.BaseResponse, err error) {
	data, err := ctrl.service.Register(body)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage("Error"), nil
	}
	if data.ID == 0 {
		return dto.DefaultErrorBaseResponseWithMessage("Phone already used"), nil
	}
	res = dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Success",
		Data:         data,
	}
	return res, nil
}

func (ctrl Controller) CheckToken(c *gin.Context) (res dto.BaseResponse, err error) {
	authDataValue, _ := c.Get("authData")
	res = dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Success",
		Data:         authDataValue,
	}
	return res, nil
}
