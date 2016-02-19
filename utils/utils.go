// utils
package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

func Unescape(x string) interface{} {
	return template.HTML(x)
}

func ConvertMarkdownToHtml(markdown string) string {
	return string(blackfriday.MarkdownBasic([]byte(markdown)))
}

//сравнение двух массивов строк s1 и s2, возвр-ет true - если s1 и s2 идентичны  -  t+
func EqStrArray(s1 []string, s2 []string) bool {
	//	f := true
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if strings.Compare(s1[i], s2[i]) != 0 {
			return false
		}
	}
	return true
}

// сортировка массива string содержащих цифры по возрастанию - t+
func SorttoUp(s []string) []string {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			s1, _ := strconv.Atoi(SplitFileName(s[i]))
			s2, _ := strconv.Atoi(SplitFileName(s[j]))
			if s1 > s2 {
				tt := s[i]
				s[i] = s[j]
				s[j] = tt
			}
		}
	}

	return s
}

// сортировка массива string содержащих цифры по убыванию  -  t+
func SorttoDown(s []string) []string {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			s1, _ := strconv.Atoi(SplitFileName(s[i]))
			s2, _ := strconv.Atoi(SplitFileName(s[j]))
			if s1 < s2 {
				tt := s[i]
				s[i] = s[j]
				s[j] = tt
			}
		}
	}

	return s
}

// выделение имени файла из строки   t+
func SplitFileName(s string) string {
	_, sf := filepath.Split(s)
	sfn := strings.Split(sf, ".")
	if sfn != nil {
		return sfn[0]
	} else {
		return ""
	}
}

//-----------функции для работы с файлами и каталогами

//возвращает список имен файлов в директории dirname  -    t-
func Getlistfileindirectory(dirname string) []string {
	listfile := make([]string, 0)
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			//fmt.Println(fi.Name(), fi.Size(), "bytes")
			listfile = append(listfile, fi.Name())
		}
	}
	return listfile
}

// сохранить в новый файл
func SaveNewstrtofile(namef string, str string) int {
	file, err := os.Create(namef)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

// сохранить файл
func Savestrtofile(namef string, str string) int {
	file, err := os.OpenFile(namef, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0776)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

//// чтение файла с именем namefи возвращение содержимое файла, иначе текст ошибки
func Readfiletxt(namef string) string {
	file, err := os.Open(namef)
	if err != nil {
		return "handle the error here"
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "error here"
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "error here"
	}
	return string(bs)
}

//-----------END функции для работы с файлами и каталогами

// в виде строки текущая дата
func GetNowDate() string {
	return (time.Now().String())[:19]
}

//  преобразование bool в строку -  t+
func Bool2String(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

//  преобразование bool в строку  - t+
func String2Bool(s string) bool {
	if s == "true" {
		return true
	} else {
		return false
	}
}

// t+
func ConcatPathFileName(namefs []string, pathstr string) []string {
	for k, v := range namefs {
		namefs[k] = pathstr + v
	}
	return namefs
}

//------------
// возвращает значение параметра params из строки str  - t+
func GetParamsFromStr(params string, str string) string {
	var val string = ""
	pos := strings.Index(str, params+":")
	if (pos == -1) || (len(params) == 0) {
		return ""
	}
	val = DelLeftSpace(DelRigthSpace(str[pos+len(params)+1:]))
	return val
}

// возвращает  - t+
func GetParamsFromList(params []string, liststr []string) map[string]string {
	var r string
	res := make(map[string]string)
	for _, z := range params {
		for _, v := range liststr {
			r = GetParamsFromStr(z, v)
			if r != "" {
				res[z] = r
				r = ""
			}
		}
	}
	return res
}

//парсинг конфиг файла map[ключ] значение_ключа  - t+
func ParseCfgFile(params []string, namef string) map[string]string {
	//	res := make(map[string]string, 0)
	str := Readfiletxt(namef)
	if len(str) == 0 {
		return nil
	}
	liststr := strings.Split(str, "\n")
	res := GetParamsFromList(params, liststr)
	return res
}

//------------
// удаление пробелов слева в строке s  - t+
func DelLeftSpace(s string) string {
	var res string = ""
	pos := len(s)
	if len(s) == 0 {
		return s
	}
	for k, v := range s {
		if v != ' ' {
			pos = k
			break
		}
	}
	res = s[pos:]
	return res
}

// удаление пробелов справа в строке s   -  t+
func DelRigthSpace(s string) string {
	var res string = ""
	res = ReverseStr(s)
	res = DelLeftSpace(res)
	res = ReverseStr(res)
	return res
}

// реверс строки s   - t+
func ReverseStr(s string) string {
	res := make([]rune, 0)
	ss := []rune(s)
	for i := len(ss) - 1; i >= 0; i-- {
		res = append(res, ss[i])
	}

	return string(res)
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
