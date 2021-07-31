package interfaces

import (
	"github.com/labstack/echo"
)

// EchoRestHandler delivery factory for echo handler
type EchoRestHandler interface {
	Mount(group *echo.Group)
}
