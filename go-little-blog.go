package main

import (
	//	"flag"
	"fmt"
	//	"html/template"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"

	//	"image"
	//	_ "image/gif"
	//	"image/jpeg"
	//	_ "image/png"
)

////------------ Объявление типов и глобальных переменных

var (
	pathposts    string // папка в которой нах-ся посты блога
	pathhtml     string // папка в которой нах-ся обычные html страницы блога
	pathtemplate string //  папка в которой нах-ся шаблоны
	kolpost      int    // кол-во постов (сообщений) на главной странице блога
	tekpost      int    // номер сообщения с которого начинается сообщения на странице
)

// структура поста в блоге
type Post struct {
	Id          string
	Title       string
	ContentText string
}

//данные для генерации страницы html
type PagePost struct {
	TitlePage string
	Posts     []Post
	Postleft  int // кол-во сообщений влево , т.е. более поздние
	Postright int // кол-во сообщений вправо , т.е. более ранние
}

func (p *Post) Print() {
	fmt.Println("Id post: ", p.Id)
	fmt.Println("Title post: ", p.Title)
	fmt.Println("ContentText post: ", p.ContentText)
}

func Print(p []Post) {
	for _, v := range p {
		v.Print()
	}
}

//------------ END Объявление типов и глобальных переменных

func unescape(x string) interface{} {
	return template.HTML(x)
}

func ConvertMarkdownToHtml(markdown string) string {
	return string(blackfriday.MarkdownBasic([]byte(markdown)))
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
func readfiletxt(namef string) string {
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

// полчение текста поста блога из файла : первая строка это заголовок сообщения, вторая и последующие это само сообщение
func GetPostfromFile(namef string) Post {
	str := readfiletxt(namef)

	sposts := strings.Index(str, "\n") // поиск первой строки - заголовка сообщения
	stitle := ""
	scontent := ""
	if sposts != -1 {
		stitle = str[0:sposts]
		if sposts+1 <= len(str) {
			scontent = str[sposts+1:]
		}
	}

	//	fmt.Println(ConvertMarkdownToHtml(scontent))

	res := Post{Id: namef, Title: ConvertMarkdownToHtml(stitle), ContentText: ConvertMarkdownToHtml(scontent)}
	//	res := Post{Id: namef, Title: stitle, ContentText: scontent}
	return res
}

func indexHandler(rr render.Render, w http.ResponseWriter, r *http.Request) {
	p := make([]Post, 0)
	namefs := Getlistfileindirectory(pathposts)
	tnamefs := namefs
	vsegopost := len(namefs)
	if len(namefs) != 0 {
		namefs = SorttoDown(namefs)
		if kolpost > len(namefs) {
			tnamefs = namefs[:]
		} else {
			tnamefs = namefs[:kolpost]
		}

		for _, namef := range tnamefs {
			p = append(p, GetPostfromFile(pathposts+string(os.PathSeparator)+namef))
		}
	} else {
		p = append(p, Post{Id: "ПОСТОВ НЕТ", Title: "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)", ContentText: ""})
	}

	rr.HTML(200, "index", &PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postright: vsegopost - kolpost})
}

// посты блога
func PostsHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
}

// посты блога
func HtmlHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	namefs := Getlistfileindirectory(pathhtml)
	if len(namefs) != 0 {
		for _, v := range namefs {
			if v == params["namepage"] {
				// ? сделать чтобы загружался страница из папки pathhtmlts
				//				rr.HTML(200, "news")
			}
		}
	}
	rr.Redirect("/")
}

// просмотр посты блога
func ViewHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	p := make([]Post, 0)
	numpost, _ := strconv.Atoi(params["numpost"])
	namefs := Getlistfileindirectory(pathposts)
	namefs = SorttoUp(namefs)
	tnamefs := namefs
	vsegopost := len(namefs)

	if numpost <= 0 {
		numpost = len(namefs)
	}

	if numpost >= len(namefs) {
		if kolpost > len(namefs) {
			tnamefs = namefs[:]
		} else {
			tnamefs = namefs[len(namefs)-kolpost:]
		}

		if len(namefs) != 0 {
			for k := len(tnamefs) - 1; k >= 0; k-- {
				namef := tnamefs[k]
				p = append(p, GetPostfromFile(pathposts+string(os.PathSeparator)+namef))
			}
		} else {
			p = append(p, Post{Id: "ПОСТОВ НЕТ", Title: "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)", ContentText: ""})
		}
		rr.HTML(200, "view", &PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postright: vsegopost - kolpost})
		return
	}
	if len(namefs) != 0 {
		kk := numpost - kolpost
		if kk <= 0 {
			kk = 0
		} else {
			kk = numpost - kolpost // - 1
		}
		tnamefs := namefs[kk:numpost]

		if len(namefs) != 0 {
			for k := len(tnamefs) - 1; k >= 0; k-- {
				namef := tnamefs[k]
				p = append(p, GetPostfromFile(pathposts+string(os.PathSeparator)+namef))
			}
		}
	} else {
		rr.Redirect("/")
		return
	}
	rr.HTML(200, "view", &PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postleft: numpost + kolpost, Postright: numpost - kolpost})
}

func main() {
	parports := ""
	m := martini.Classic()

	//	martini.Env = martini.Prod

	if martini.Env == martini.Prod {
		parports = "80"
	} else {
		parports = "1111"
	}

	//	if !parse_args() {
	//		return
	//	}

	//--------параметры программы------
	pathposts = "posts"
	pathhtml = "html"
	//	pathtemplate = "templates"
	pathtemplate = "templates\\uno-theme"
	kolpost = 3 // кол-во постов которые видны на странице
	//--------------

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Use(render.Renderer(render.Options{
		Directory:  pathtemplate,                        // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Extensions: []string{".tmpl", ".html"}}))

	m.Get("/", indexHandler)
	m.Get("/html/:namepage", HtmlHandler)
	m.Get("/view/:numpost", ViewHandler)
	m.RunOnAddr(":" + parports)

}

//// функция парсинга аргументов программы
//func parse_args() bool {
//	flag.StringVar(&hd, "hd", "", "Рабочая папка где нах-ся папки пользователей для сохранения ")
//	flag.StringVar(&user, "user", "", "Рабочая папка где нах-ся папки пользователей для сохранения ")
//	flag.Parse()
//	pathcfg = hd
//	if user == "" {
//		tekuser = "testuser"
//	} else {
//		tekuser = user
//	}
//	return true
//}
