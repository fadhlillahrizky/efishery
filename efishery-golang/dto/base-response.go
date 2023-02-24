package dto

type BaseResponse struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
}

func DefaultErrorBaseResponse(err error) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}
}

func DefaultBaseResponseWithError(err error) (BaseResponse, error) {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}, err
}

func DefaultErrorBaseResponseWithMessage(message string) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      message,
	}
}

func DefaultSuccessResponseWithMessage(msg string) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      true,
		MessageTitle: "",
		Message:      msg,
	}
}
