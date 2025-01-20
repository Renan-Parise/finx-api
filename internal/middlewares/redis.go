package middlewares

import (
	"bytes"
	"context"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/Renan-Parise/finances/internal/errors"
	"github.com/Renan-Parise/finances/internal/redis"
	"github.com/gin-gonic/gin"
)

const cacheDuration = 2 * time.Minute

func RedisCacheMiddleware(c *gin.Context) {
	ctx := context.Background()

	cacheKey := generateCacheKey(c.Request)

	cached, err := redis.Get(ctx, cacheKey)
	if err == nil && cached != nil {
		c.Data(http.StatusOK, "application/json", cached)
		c.Abort()
		return
	}

	writer := c.Writer
	buff := &responseBuffer{body: bytes.NewBufferString("")}
	c.Writer = &ginResponseWriter{
		ResponseWriter: writer,
		body:           buff.body,
	}

	c.Next()

	statusCode := c.Writer.Status()
	if statusCode == http.StatusOK {
		err := redis.Set(ctx, cacheKey, buff.body.Bytes(), cacheDuration)
		if err != nil {
			errors.NewServiceError("Failed to cache response. Error: %v" + err.Error())
		}
	}

	c.Writer = writer
}

func generateCacheKey(req *http.Request) string {
	data := req.URL.Path + "?" + req.URL.RawQuery
	h := fnv.New64a()
	h.Write([]byte(data))
	return "cache:" + string(h.Sum(nil))
}

type responseBuffer struct {
	body *bytes.Buffer
}

type ginResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *ginResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *ginResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
