package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomNumber int
var randomString string

//ランダムな英文字列を生成
func randomMaker() {
	rand.Seed(time.Now().UnixNano())
	const baseString = "abcdefghijklmnopqrstuvwxyz"
	var letterNumber int = rand.Intn(100)
	var letter []string

	for i := 0; i < letterNumber; i++ {
		letter = append(letter, string(baseString[rand.Intn(26)]))
	}
	randomString := strings.Join(letter, "")
	randomNumber := letterNumber
	fmt.Println(randomNumber)
	fmt.Println(randomString)
}
