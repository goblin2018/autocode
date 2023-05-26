package ctx

import (
	"github.com/gin-gonic/gin"
)

// Custon Context data and funcs
type Context struct {
	*gin.Context
}
