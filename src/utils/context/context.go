package context

import (
	"bm-support/config/database"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Context struct {
	Ctx        *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	UserID     uint64
	Token      string
}

// NewContext Create new instance
func NewContext(c *gin.Context) Context {
	ctx := context()
	return ctx(c)
}

type Payload struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Payload interface{} `json:"payload"`
}

// context init context
func context() func(c *gin.Context) Context {
	return func(c *gin.Context) Context {
		var context Context

		// Init gin Context
		context.Ctx = c

		// Connect to DB
		context.DB = database.OpenConnectionToDB()

		// Return Connection
		context.Connection = database.ReturnConnection(context.DB)

		// Centrlize authentication UserID
		context.UserID, _ = strconv.ParseUint(context.Ctx.Request.Header.Get("user"), 10, 64)

		return context
	}
}

// destroy close connection
func (ctx Context) Destroy() {
	ctx.Connection.Close()
}
