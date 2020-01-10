/*
 * 登录服务器
 * 处理用户登录请求，去mysql或者redis
 * 验证用户，将token和用户信息放入redis
 * 设置redis过期时间
 */

package main

import (
	"net/http"
	//"encoding/json"
	"github.com/labstack/echo"
)

func main() {
	//创建echo实例
	e := echo.New()

	//添加test测试接口和响应
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "登录服务测试正常!")
	})

	//启动服务
	e.Logger.Fatal(e.Start(":8090"))
}
