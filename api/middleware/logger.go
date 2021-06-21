package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

func HttpLoggerMiddleware(c *gin.Context) {
	start := time.Now()
	f, err := os.OpenFile("storage/logs/log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = fmt.Fprintln(f, c.Request.Method+" "+c.Request.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	c.Next()

	statusCode := strconv.Itoa(c.Writer.Status())
	_, err = fmt.Fprintln(
		f,
		c.Request.RequestURI+" "+statusCode+" duration -> "+time.Since(start).String(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
