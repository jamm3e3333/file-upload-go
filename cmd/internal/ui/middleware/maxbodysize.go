package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaxBodySizeBytes 4 MiB
const MaxBodySizeBytes = 4 << 20

func CheckMaxBodySize(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MaxBodySizeBytes)

	c.Next()
}
