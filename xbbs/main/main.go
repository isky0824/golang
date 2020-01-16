/*
 * xbbs主程序
 */

package main

import (
	_ "encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//echo 路由
	/*
	 *e.POST("/users", saveUser)
	 *e.GET("/users/:id", getUser)
	 *e.PUT("/users/:id", updateUser)
	 *e.DELETE("/users/:id", deleteUser)
	 */

	//创建echo实例
	e := echo.New()

	//添加/接口和响应
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello go echo!")
	})

	//添加/userid接口和响应
	e.GET("/userid/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	//添加/username接口和响应
	e.GET("/username", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, name)
	})

	// e.POST("/save", save)
	e.POST("/save", func(c echo.Context) error {
		// Get name and email
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:"+name+", email:"+email)
	})

	//启动服务
	e.Logger.Fatal(e.Start(":8090"))
}
