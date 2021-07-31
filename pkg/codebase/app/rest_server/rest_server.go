package restserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo"
	echoMidd "github.com/labstack/echo/middleware"

	"webarticles/pkg/codebase/factory"
	"webarticles/pkg/config"
	"webarticles/pkg/logger"
	"webarticles/pkg/wrapper"
)

type restServer struct {
	serverEngine *echo.Echo
	service      factory.ServiceFactory
}

// NewServer create new REST server
func NewServer(service factory.ServiceFactory) factory.AppServerFactory {
	server := &restServer{
		serverEngine: echo.New(),
		service:      service,
	}

	return server
}

func (h *restServer) Serve() {
	h.serverEngine.HTTPErrorHandler = wrapper.CustomHTTPErrorHandler
	h.serverEngine.Use(echoMidd.CORS())
	h.serverEngine.Use(echoMidd.Recover())

	h.serverEngine.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":   fmt.Sprintf("Service %s up and running", h.service.Name()),
			"timestamp": time.Now().Format(time.RFC3339Nano),
		})
	})

	h.serverEngine.GET("/application/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":   fmt.Sprintf("Service %s up and running", h.service.Name()),
			"timestamp": time.Now().Format(time.RFC3339Nano),
		})
	})

	if config.BaseEnv().UseSwagger {
		h.serverEngine.File("/docs/swagger", "docs/swagger/docs.json")
		h.serverEngine.Static("/docs", "docs/swagger")
	}

	restRootPath := h.serverEngine.Group("",
		EchoRestTracerMiddleware, echoMidd.Logger(),
	)
	for _, m := range h.service.GetModules() {
		if h := m.RestHandler(); h != nil {
			h.Mount(restRootPath)
		}
	}

	var routes strings.Builder
	httpRoutes := h.serverEngine.Routes()
	sort.Slice(httpRoutes, func(i, j int) bool {
		return httpRoutes[i].Path < httpRoutes[j].Path
	})
	for _, route := range httpRoutes {
		if !strings.Contains(route.Name, "(*Group)") {
			routes.WriteString(logger.StringGreen(fmt.Sprintf("[REST-ROUTE] %-6s %-30s --> %s\n", route.Method, route.Path, route.Name)))
		}
	}
	fmt.Print(routes.String())

	h.serverEngine.HideBanner = true
	h.serverEngine.HidePort = true
	port := fmt.Sprintf(":%d", config.BaseEnv().RESTPort)
	fmt.Printf("\x1b[34;1m⇨ REST server run at port [::]%s\x1b[0m\n\n", port)
	if err := h.serverEngine.Start(port); err != nil {
		switch e := err.(type) {
		case *net.OpError:
			panic(e)
		}
	}
}

func (h *restServer) Shutdown(ctx context.Context) {
	deferFunc := logger.LogWithDefer("Stopping REST HTTP server...")
	defer deferFunc()

	if err := h.serverEngine.Shutdown(ctx); err != nil {
		panic(err)
	}
}
