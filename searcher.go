package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/user", show)
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))

	//post
	e.POST("/users", postData)

}
