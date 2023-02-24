package auth

import (
	"efishery-golang/middleware"
	"efishery-golang/repository"
	"efishery-golang/utils/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl Controller
}

func NewRequestHandler(
	DB *gorm.DB,
) *RequestHandler {
	return &RequestHandler{
		ctrl: NewController(
			NewService(
				repository.NewUser(DB),
			),
		),
	}
}

func (h RequestHandler) Handle(router *gin.RouterGroup) {
	router.POST(
		"/login",
		h.login,
	)
	router.POST(
		"/register",
		h.register,
	)
	router.GET(
		"/check-token",
		middleware.Authenticate(),
		h.CheckToken,
	)
}

func (h RequestHandler) login(c *gin.Context) {
	var req LoginRequest
	if !validator.BindAndValidateWithAbort(c, &req) {
		return
	}

	res, err := h.ctrl.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h RequestHandler) register(c *gin.Context) {
	var req RegisterRequest
	if !validator.BindAndValidateWithAbort(c, &req) {
		return
	}

	res, err := h.ctrl.register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h RequestHandler) CheckToken(c *gin.Context) {
	res, _ := h.ctrl.CheckToken(c)
	c.JSON(http.StatusOK, res)
	return
}
