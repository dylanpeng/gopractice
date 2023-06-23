package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	// 随机数
	ran = rand.New(rand.NewSource(time.Now().UnixMilli()))
	// 请求处理时间指标
	httpRequestDurationVec *prometheus.HistogramVec
	// 请求个数指标
	httpRequestCountVec *prometheus.CounterVec
	// DefaultBuckets prometheus buckets in seconds.
	DefaultBuckets = []float64{0.1, 0.3, 0.5, 1.0, 3.0, 5.0}
)

func main() {
	router := gin.Default()
	router.Use(Middleware)

	router.GET("/metrics", Metrics())
	router.GET("/favicon.ico", CommonControl)

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/A", CommonControl)
		apiGroup.GET("/B", CommonControl)
		apiGroup.GET("/C", CommonControl)
		apiGroup.GET("/D", CommonControl)
		apiGroup.GET("/E", CommonControl)
	}

	_ = router.Run("0.0.0.0:15000")
}

// 初始化prometheus
func init() {
	httpRequestDurationVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_server_requests_seconds",
		Help:    "How long it took to process the HTTP request, partitioned by status code, method and HTTP path.",
		Buckets: DefaultBuckets,
	}, []string{"code", "method", "uri"})

	httpRequestCountVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_server_requests_count",
		Help: "How long it took to process the HTTP request, partitioned by status code, method and HTTP path.",
	}, []string{"code", "method", "uri"})

	prometheus.MustRegister(httpRequestDurationVec, httpRequestCountVec)
}

// 统一处理函数
func CommonControl(ctx *gin.Context) {
	time.Sleep(time.Duration(ran.Intn(100)) * time.Millisecond)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// 中间件记录请求指标
func Middleware(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()

	// 返回状态码
	statusCode := strconv.Itoa(ctx.Writer.Status())
	// 标签
	labels := []string{statusCode, ctx.Request.Method, ctx.Request.URL.Path}
	// 请求处理时间
	duration := float64(time.Since(start).Nanoseconds()) / 1000000

	// 添加指标
	httpRequestDurationVec.WithLabelValues(labels...).Observe(duration)
	httpRequestCountVec.WithLabelValues(labels...).Inc()
	fmt.Printf("code: %d | method: %s | path: %s | duration: %f\n ", statusCode, ctx.Request.Method, ctx.Request.URL.Path, duration)
}

// http metrics 指标页面，配置给prometheus
func Metrics() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
