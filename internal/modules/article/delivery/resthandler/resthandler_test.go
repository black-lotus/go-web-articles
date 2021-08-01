package resthandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webarticles/internal/modules/article/domain"
	"webarticles/internal/modules/article/usecase"
	"webarticles/pkg/config"
	"webarticles/pkg/testhelper"
	"webarticles/pkg/validator"
	"webarticles/pkg/wrapper"

	"github.com/brianvoe/gofakeit"
	"github.com/integralist/go-findroot/find"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ctx             context.Context
	restHandlerMock *RestHandler
	usecaseMock     *usecase.MockArticleUsecase
)

func doMock() {
	// set usecase
	usecaseMock = &usecase.MockArticleUsecase{}

	// set root
	root, _ := find.Repo()

	env := config.Env{}
	env.JSONSchemaDir = fmt.Sprintf("%s/api/jsonschema", root.Path)
	config.SetEnv(env)

	// set json validator
	jsonValidator := validator.NewValidator()

	// set rest handler
	restHandlerMock = &RestHandler{
		validator:      jsonValidator,
		articleUsecase: usecaseMock,
	}
}

func TestNewRestHandler(t *testing.T) {
	testName := testhelper.SetTestcaseName(1, "new rest handler")

	t.Run(testName, func(t *testing.T) {
		doMock()

		NewRestHandler(restHandlerMock.validator, restHandlerMock.articleUsecase)
	})
}

func TestRestHandler_Mount(t *testing.T) {
	testName := testhelper.SetTestcaseName(1, "rest handler mount")

	t.Run(testName, func(t *testing.T) {
		doMock()

		// set rest handler
		restHandler := restHandlerMock

		// set echo
		echoHandler := echo.New()
		groupEcho := echoHandler.Group(gofakeit.Word())

		restHandler.Mount(groupEcho)
	})
}

func TestRestHandler_findAll(t *testing.T) {
	type findAllResult struct {
		articles []*domain.Article
		err      error
	}

	testCase := map[string]struct {
		findAll *findAllResult
		query   string
		assert  func(*testing.T, echo.Context, *wrapper.HTTPResponse)
	}{
		testhelper.SetTestcaseName(1, "Given happy case"): {
			findAll: &findAllResult{
				articles: []*domain.Article{},
				err:      nil,
			},
			query: "query=some-query&author=dondon",
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "SUCCESS", response.Code)
				require.NotNil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(2, "Given usecase return error"): {
			findAll: &findAllResult{
				err: fmt.Errorf("Something unexpected happens"),
			},
			query: "query=some-query&author=dondon",
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				require.Nil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
	}
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			// set HTTP mock
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles?%s", test.query), strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			doMock()

			if test.findAll != nil {
				usecaseMock.On("FindAll", mock.Anything, mock.Anything).Return(test.findAll.articles, test.findAll.err).Once()
			}

			// set rest handler
			restHandler := restHandlerMock

			// set handler
			err := restHandler.findAll(c)

			// validate result
			require.NoError(t, err)
			content := new(wrapper.HTTPResponse)
			err = json.Unmarshal(rec.Body.Bytes(), content)
			require.NoError(t, err)

			// do assertion
			test.assert(t, c, content)
		})
	}
}

func TestRestHandler_create(t *testing.T) {
	article := domain.Article{
		Author: "Dondon",
		Title:  "why golang isn't proposed for dummy developer?",
		Body:   "Because they miss-understood the flexibility feature",
	}
	type createResult struct {
		article *domain.Article
		err     error
	}

	type otherData struct {
		username string
	}

	testCase := map[string]struct {
		payload      *domain.Article
		otherData    *otherData
		createResult *createResult
		assert       func(*testing.T, echo.Context, *wrapper.HTTPResponse)
	}{
		testhelper.SetTestcaseName(1, "Given happy case"): {
			payload:      &article,
			createResult: &createResult{article: &domain.Article{}},
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusCreated, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "SUCCESS", response.Code)
				require.NotNil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(2, "Given usecase error"): {
			payload:      &article,
			createResult: &createResult{err: fmt.Errorf("Something unexpected happens")},
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "Something unexpected happens", response.Data)
				assert.Equal(t, "Failed to create data", response.Message)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(3, "Given value of the bodies are empty"): {
			payload: &domain.Article{},
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				//cannot be asserted since the ordering was random
				//assert.Equal(t, "author: String length must be greater than or equal to 3\ntitle: String length must be greater than or equal to 3\nbody: String length must be greater than or equal to 3", response.Data)
				require.NotNil(t, response.Data)
				assert.Equal(t, "Bad Request", response.Message)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(4, "Given wrong payload"): {
			otherData: &otherData{
				username: "dondon",
			},
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "Bad Request", response.Message)
				require.NotNil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
	}
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			// set payload
			var payload string
			if test.payload != nil {
				payload = testhelper.CreateHttpRequestBodyMock(test.payload)
			}
			if test.otherData != nil {
				payload = testhelper.CreateHttpRequestBodyMock(test.otherData)
			}

			// set http mock
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			doMock()

			if test.createResult != nil {
				usecaseMock.On("CreateArticle", mock.Anything, mock.Anything).Return(test.createResult.article, test.createResult.err).Once()
			}

			// set rest handler
			restHandler := restHandlerMock

			// set handler
			err := restHandler.create(c)

			// validate result
			require.NoError(t, err)
			content := new(wrapper.HTTPResponse)
			err = json.Unmarshal(rec.Body.Bytes(), content)
			require.NoError(t, err)

			// do assertion
			test.assert(t, c, content)
		})
	}
}

func TestRestHandler_findByID(t *testing.T) {
	type findByID struct {
		article *domain.Article
		err     error
	}

	testCase := map[string]struct {
		findByID *findByID
		ID       string
		assert   func(*testing.T, echo.Context, *wrapper.HTTPResponse)
	}{
		testhelper.SetTestcaseName(1, "Given happy case"): {
			findByID: &findByID{article: &domain.Article{}},
			ID:       "1",
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "SUCCESS", response.Code)
				require.NotNil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(2, "Given usecase error"): {
			findByID: &findByID{err: fmt.Errorf("Something unexpected happens")},
			ID:       "1",
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				assert.Equal(t, "Something unexpected happens", response.Data)
				assert.Equal(t, "Failed to find data", response.Message)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(3, "Given empty param (ID)"): {
			ID: "",
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				require.Nil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
		testhelper.SetTestcaseName(4, "Given invalid param (ID)"): {
			ID: gofakeit.Word(),
			assert: func(t *testing.T, c echo.Context, response *wrapper.HTTPResponse) {
				resp := c.Response()
				assert.Equal(t, http.StatusBadRequest, resp.Status)
				assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
				require.Nil(t, response.Data)

				usecaseMock.AssertExpectations(t)
			},
		},
	}

	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			// set http mock
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/articles/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("id")
			c.SetParamValues(test.ID)

			doMock()

			if test.findByID != nil {
				usecaseMock.On("FindByID", mock.Anything, mock.Anything).Return(test.findByID.article, test.findByID.err).Once()
			}

			// set rest handler
			restHandler := restHandlerMock

			// set handler
			err := restHandler.findByID(c)

			// validate result
			require.NoError(t, err)
			content := new(wrapper.HTTPResponse)
			err = json.Unmarshal(rec.Body.Bytes(), content)
			require.NoError(t, err)

			// do assertion
			test.assert(t, c, content)
		})
	}
}
