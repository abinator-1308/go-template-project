package server

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/errors"
	"github.com/devlibx/gox-base/metrics"
	"github.com/gin-gonic/gin"
	"github.com/uber-go/tally"
	"net/http"
	"sync"
)

// GinContextToContextMiddleware will add the Gin context in context so we can use it in request
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GinContextFromHttpRequestVerified this will ignore the error, this assumes we know this request has gone via
// Gin MW and has the Gin context
func GinContextFromHttpRequestVerified(r *http.Request) *gin.Context {
	gc, err := GinContextFromHttpRequest(r)
	if err != nil {
		fmt.Println("********** Unexpected error **********", err)
	}
	return gc
}

// GinContextFromHttpRequest will get the gin context from Http Request
func GinContextFromHttpRequest(r *http.Request) (*gin.Context, error) {
	return GinContextFromContext(r.Context())
}

// GinContextFromContext will get the gin context any context Http Request
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		return nil, errors.New("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}

// EnsureGinContextWrapper is common helper to make sure we have Gin context
func EnsureGinContextWrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Ensure we have gin context in request
		_, err := GinContextFromHttpRequest(request)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

		h.ServeHTTP(writer, request)
	}
}

func MetricWrapper(h http.HandlerFunc, cf gox.CrossFunction, metricName string) http.HandlerFunc {
	initOnce := sync.Once{}
	var callCounter metrics.Counter
	var hist metrics.Histogram
	return func(w http.ResponseWriter, r *http.Request) {
		initOnce.Do(func() {
			callCounter = cf.Metric().Counter("handler__" + metricName)
			hist = cf.Metric().Histogram("hist__handler__"+metricName, tally.ValueBuckets{0.5, 0.9, 0.99, 0.999, 1})
		})
		callCounter.Inc(1)
		sw := hist.Start()
		defer sw.Stop()

		h.ServeHTTP(w, r)
	}
}
