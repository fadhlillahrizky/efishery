package resources

import (
	"efishery-golang/middleware"
	"efishery-golang/repository"
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
				repository.NewResources(DB),
			),
		),
	}
}

func (h RequestHandler) Handle(router *gin.RouterGroup) {
	router.GET(
		"/",
		middleware.Authenticate(),
		h.List,
	)
}

func (h RequestHandler) List(c *gin.Context) {

	res, err := h.ctrl.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
