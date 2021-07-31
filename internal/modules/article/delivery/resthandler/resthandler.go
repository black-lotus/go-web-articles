package resthandler

import (
	"fmt"
	config "webarticles/configs"
	"webarticles/internal/modules/article/usecase"

	"github.com/labstack/echo"
)

// RestHandler handler
type RestHandler struct {
	articleUsecase usecase.ArticleUsecase
}

// NewRestHandler create new rest handler
func NewRestHandler(articleUsecase usecase.ArticleUsecase) *RestHandler {
	return &RestHandler{
		articleUsecase: articleUsecase,
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	g := fmt.Sprintf("/%s", config.GetEnv().ServicePath)
	gRoot := root.Group(g)

	datasources := gRoot.Group("/article")
	datasources.GET("", h.findAll)
	datasources.POST("", h.create)
	datasources.GET("/:id", h.findByID)
}

func (h *RestHandler) findAll(c echo.Context) error {
	return nil
}

func (h *RestHandler) create(c echo.Context) error {
	return nil
}

func (h *RestHandler) findByID(c echo.Context) error {
	return nil
}
