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

// сортировка массива string
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

// сортировка массива string
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

// выделение имени файла из строки
func SplitFileName(s string) string {
	_, sf := filepath.Split(s)
	sfn := strings.Split(sf, ".")
	if sfn != nil {
		return sfn[0]
	} else {
		return ""
	}

}

//возвращает список имен файлов в директории dirname
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

// в виде строки текущая дата
func GetNowDate() string {
	return (time.Now().String())[:19]
}

//  преобразование bool в строку
func Bool2String(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

//  преобразование bool в строку
func String2Bool(s string) bool {
	if s == "true" {
		return true
	} else {
		return false
	}
}
