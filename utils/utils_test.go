package utils

import (
	//	"fmt"
	"strings"
	"testing"
)

//---------тест для конфигурационного файла

// возвращает значение параметра params из строки str
//func GetParamsFromStr(params string, str string) string
func TestGetParamsFromStr(t *testing.T) {
	var (
		par string = "path"
		str string = "path: c:\\oilnur\\  "
	)
	res := GetParamsFromStr(par, str)
	if res != "c:\\oilnur\\" {
		t.Fatalf("неправильный результат", res)
	}
	//---
	str = "path:oilnur sysadmin qq  "
	res = GetParamsFromStr(par, str)
	if res != "oilnur sysadmin qq" {
		t.Fatalf("неправильный результат", res)
	}
	//---
	par = "pathposts"
	str = "pathtemplate: uno-theme"
	res = GetParamsFromStr(par, str)
	if res != "" {
		t.Fatalf("неправильный результат", res)
	}
}

//GetParamsFromList(params []string, liststr []string) map[string]string
func TestGetParamsFromList(t *testing.T) {
	var (
		params      = []string{"pathposts", "kolpost", "pathtemplate"}
		liststr     = []string{"pathposts: posts", "pathhtml: html", "pathtemplate: uno-theme", "kolpost: 3"}
		liststrnull = []string{}
	)

	s := GetParamsFromList(params, liststr)
	if (s[params[0]] != "posts") && (s[params[1]] != "3") && (s[params[2]] != "uno-theme") {
		t.Fatalf("неправильный результат", s)
	}
	s = GetParamsFromList(params, liststrnull)
	if len(s) != 0 {
		t.Fatalf("неправильный результат ", s)
	}
}

//парсинг конфиг файла map[ключ] значение_ключа
//func ParseCfgFile(namef string) map[string]string
func TestParseCfgFile(t *testing.T) {
	var (
		params = []string{"pathposts", "kolpost", "pathtemplate"}
	)
	s := ParseCfgFile(params, "config.cfg")
	if (s[params[0]] != "posts") && (s[params[1]] != "3") && (s[params[2]] != "uno-theme") {
		t.Fatalf("неправильный результат", s)
	}
	s = ParseCfgFile(params, "configs.cfg")
	if len(s) != 0 {
		t.Fatalf("неправильный результат ", s)
	}
}

//---------END тест для конфигурационного файла

func TestDelLeftSpace(t *testing.T) {
	var (
		str string
		res string
	)
	str = "  posts   "
	res = DelLeftSpace(str)
	if res != "posts   " {
		t.Fatalf("неправильный результат ", res)
	}
	res = DelLeftSpace("")
	if len(res) != 0 {
		t.Fatalf("неправильный результат ", res)
	}
	res = DelLeftSpace("    ")
	if len(res) != 0 {
		t.Fatalf("неправильный результат ", res)
	}
}

func TestDelRigthSpace(t *testing.T) {
	var (
		str string
		res string
	)
	str = "  posts   "
	res = DelRigthSpace(str)
	if res != "  posts" {
		t.Fatalf("неправильный результат ", res)
	}
	res = DelRigthSpace("")
	if len(res) != 0 {
		t.Fatalf("неправильный результат ", res)
	}
	res = DelRigthSpace("    ")
	if len(res) != 0 {
		t.Fatalf("неправильный результат ", res)
	}

	str = "  posts "
	res = DelRigthSpace(str)
	if res != "  posts" {
		t.Fatalf("неправильный результат ", res)
	}
}

// ReverseStr(s string) string
func TestReverseStr(t *testing.T) {
	s := "abcdef"
	res := ReverseStr(s)
	if res != "fedcba" {
		t.Fatalf("неправильный результат ", res)
	}
	s = "  abcdef "
	res = ReverseStr(s)
	if res != " fedcba  " {
		t.Fatalf("неправильный результат ", res)
	}
	s = ""
	res = ReverseStr(s)
	if res != "" {
		t.Fatalf("неправильный результат ", res)
	}

}

//сравнение двух массивов строк s1 и s2, возвр-ет true - если s1 и s2 идентичны  -  t-
//func EqStrArray(s1 []string, s2 []string) bool {
func TestEqStrArray(t *testing.T) {
	s1 := []string{"10", "7", "8", "4", "0"}
	s2 := []string{"0", "4", "7", "8", "10"}
	res := EqStrArray(s1, s2)
	if res {
		t.Fatalf("неправильный результат ", res)
	}

	res = EqStrArray(s1, s1)
	if !res {
		t.Fatalf("неправильный результат ", res)
	}

	s3 := make([]string, 0)
	res = EqStrArray(s1, s3)
	if res {
		t.Fatalf("неправильный результат ", res)
	}

	res = EqStrArray(s3, s3)
	if !res {
		t.Fatalf("неправильный результат ", res)
	}
}

// сортировка массива string содержащих цифры
//func SorttoUp(s []string) []string {
func TestSorttoUp(t *testing.T) {
	s := []string{"10", "7", "8", "4", "0"}
	sres := []string{"0", "4", "7", "8", "10"}
	res := SorttoUp(s)
	if !EqStrArray(res, sres) {
		t.Fatalf("неправильный результат ", res)
	}

	s = []string{"ф", "7", "8", "4", "0"}
	sres = []string{"ф", "0", "4", "7", "8"}
	res = SorttoUp(s)
	if !EqStrArray(res, sres) {
		t.Fatalf("неправильный результат ", res)
	}
}

// сортировка массива string содержащих цифры по убыванию
//func SorttoDown(s []string) []string {
func TestSorttoDown(t *testing.T) {
	s := []string{"10", "7", "8", "4", "0"}
	sres := []string{"10", "8", "7", "4", "0"}
	res := SorttoDown(s)
	if !EqStrArray(res, sres) {
		t.Fatalf("неправильный результат ", res)
	}

	//	s = []string{"ф", "7", "8", "4", "0"}
	//	sres = []string{"ф", "8", "7", "4", "0"}
	//	res = SorttoDown(s)
	//	if !EqStrArray(res, sres) {
	//		t.Fatalf("неправильный результат ", res)
	//	}
}

// выделение имени файла из строки
//func SplitFileName(s string) string {
func TestSplitFileName(t *testing.T) {
	s := "filename1.txt"
	sres := "filename1"
	res := SplitFileName(s)
	if strings.Compare(res, sres) != 0 {
		t.Fatalf("неправильный результат ", res)
	}

	s = "привет1.txt"
	sres = "привет1"
	res = SplitFileName(s)
	if strings.Compare(res, sres) != 0 {
		t.Fatalf("неправильный результат ", res)
	}

	s = ""
	sres = ""
	res = SplitFileName(s)
	if strings.Compare(res, sres) != 0 {
		t.Fatalf("неправильный результат ", res)
	}

	s = "c:\\dir\\filename1.txt"
	sres = "filename1"
	res = SplitFileName(s)
	if strings.Compare(res, sres) != 0 {
		t.Fatalf("неправильный результат ", res)
	}

}
