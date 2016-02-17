// go-little-blog_test
package main

import (
	"fmt"
	//	"os"
	"testing"

	//	"go-little-blog/utils"
)

func TestParseCfgFile(t *testing.T) {
	fmt.Println(ParseCfgFile("config.cfg"))
}

// возвращает значение параметра params из строки str
//func GetParamsFromStr(params string, str string) string {
func TestGetParamsFromStr(t *testing.T) {
	var (
		par string = "path"
		str string = "path: c:\\oilnur\\  "
	)
	res := GetParamsFromStr(par, str)
	if res != "c:\\oilnur\\" {
		t.Fatalf("неправильный результат", res)
	}
}
