package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func show(c echo.Context) error {
	fmt.Println("done")
	name := c.QueryParam("name")
	email := c.QueryParam("email")
	id, _ := strconv.Atoi(c.QueryParam("id"))
	u := new(User)
	u.Email = email
	u.Name = name
	u.Id = id
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
	//このuをJSON形式に直す
}

func postData(c echo.Context) error {
	// requestBody := &RequestBody{
	// 	Something: "Hello",
	// }
	fmt.Println("done")
	u := new(User)
	name := c.QueryParam("name")
	email := c.QueryParam("email")
	id, _ := strconv.Atoi(c.QueryParam("id"))
	u.Email = email
	u.Name = name
	u.Id = id

	jsonString, err := json.Marshal(u)
	if err != nil {
		panic("Error")
	}

	endpoint := "http://localhost:1323/users"
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	fmt.Println(string(byteArray))
	return c.JSON(200, u)
}
