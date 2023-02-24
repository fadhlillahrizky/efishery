package resources

import (
	"efishery-golang/dto"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) Controller {
	return Controller{
		service: service,
	}
}

func (ctrl Controller) List() (res dto.BaseResponse, err error) {
	data, err := ctrl.service.List()
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage("Error"), nil
	}
	res = dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Success",
		Data:         data,
	}
	return res, nil
}
