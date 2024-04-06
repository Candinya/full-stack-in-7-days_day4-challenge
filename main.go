package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// 创建一个 echo 实例
	e := echo.New()

	// 准备中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 准备路由
	e.GET("/", handleGetPrime)

	// 开始监听服务器
	e.Logger.Fatal(e.Start(":1323"))
}

type getPrimeReq struct {
	Start *int `query:"start"`
	End   *int `query:"end"`
}

func handleGetPrime(c echo.Context) error {
	var req getPrimeReq
	if err := c.Bind(&req); err != nil {
		zap.L().Error("请求绑定失败", zap.Error(err))
		return c.String(http.StatusInternalServerError, "请求绑定失败")
	}
	if req.Start == nil || req.End == nil {
		return c.String(http.StatusBadRequest, "缺少 start 或 end")
	}

	if *req.Start > *req.End {
		req.Start, req.End = req.End, req.Start
	}

	var primeStrs []string
	for i := *req.Start; i <= *req.End; i++ {
		if isPrime(i) {
			primeStrs = append(primeStrs, strconv.Itoa(i))
		}
	}

	return c.String(http.StatusOK, strings.Join(primeStrs, "\n"))
}

// 判断是否为素数
func isPrime(num int) bool {
	if num <= 1 {
		// 不考虑负数
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}