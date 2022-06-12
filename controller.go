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

type RequestBody struct {
	Something string
}

func postData(c echo.Context) error {
	requestBody := &RequestBody{
		Something: "Hello",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}

	endpoint := "http://localhost:1323/users"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
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

	fmt.Printf("%#v", string(byteArray))
	return c.String(http.StatusOK, "OK")
}
