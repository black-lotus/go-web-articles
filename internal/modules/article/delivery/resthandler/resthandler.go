package resthandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	config "webarticles/configs"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/usecase"
	"webarticles/pkg/codebase/interfaces"
	"webarticles/pkg/http/helper"
	"webarticles/pkg/logger"
	"webarticles/pkg/wrapper"

	"github.com/labstack/echo"
)

const (
	messageBadRequest   string = "Bad Request"
	messageCreateFailed string = "Failed to create data"
	messageDataNotFound string = "Failed to find data"
	messageSuccess      string = "SUCCESS"
)

// RestHandler handler
type RestHandler struct {
	validator      interfaces.Validator
	articleUsecase usecase.ArticleUsecase
}

// NewRestHandler create new rest handler
func NewRestHandler(validator interfaces.Validator,
	articleUsecase usecase.ArticleUsecase) *RestHandler {
	return &RestHandler{
		validator:      validator,
		articleUsecase: articleUsecase,
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	g := fmt.Sprintf("/%s", config.GetEnv().ServicePath)
	gRoot := root.Group(g)

	datasources := gRoot.Group("/articles")
	datasources.GET("", h.findAll)
	datasources.POST("", h.create)
	datasources.GET("/:id", h.findByID)
}

func (h *RestHandler) findAll(c echo.Context) error {
	ctx := c.Request().Context()

	var filter domain.Filter
	if err := helper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageBadRequest, err).JSON(c.Response())
	}

	body, _ := json.Marshal(filter)
	if err := h.validator.ValidateDocument("article/find_all", body); err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageBadRequest, err.Error()).JSON(c.Response())
	}

	result, err := h.articleUsecase.FindAll(ctx, &filter)
	if err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, messageSuccess, result).JSON(c.Response())
}

func (h *RestHandler) create(c echo.Context) error {
	ctx := c.Request().Context()

	body, _ := ioutil.ReadAll(c.Request().Body)
	if err := h.validator.ValidateDocument("article/create", body); err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageBadRequest, err.Error()).JSON(c.Response())
	}

	var payload domain.Article
	if err := json.Unmarshal(body, &payload); err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	result, err := h.articleUsecase.CreateArticle(ctx, &payload)
	if err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageCreateFailed, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, messageSuccess, result).JSON(c.Response())
}

func (h *RestHandler) findByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageBadRequest).JSON(c.Response())
	}

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageBadRequest).JSON(c.Response())
	}

	result, err := h.articleUsecase.FindByID(ctx, &id)
	if err != nil {
		logger.LogRed(err.Error())
		return wrapper.NewHTTPResponse(http.StatusBadRequest, messageDataNotFound, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, messageSuccess, result).JSON(c.Response())
}
