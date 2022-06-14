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
	e.POST("/users", postData)
	//post
	// ps := url.Values{}
	// ps.Add("id", "1")
	// ps.Add("name", "moripi")
	// fmt.Println(ps.Encode())
	// res, err := http.PostForm("http://localhost:1323/users", ps)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Print(string(body))
}
