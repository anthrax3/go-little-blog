package main

import (
	//	"flag"
	"fmt"
	//	"html/template"
	"net/http"
	"os"
	//	"strconv"
	"html/template"
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
	hd   string
	user string
)

var (
	pathposts string // папка в которой нах-ся посты блога
	pathhtml  string // папка в которой нах-ся обычные html страницы блога
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

	//	namef := ""
	namefs := Getlistfileindirectory(pathposts)
	p := make([]Post, 0)
	if len(namefs) != 0 {
		for _, namef := range namefs {
			p = append(p, GetPostfromFile(pathposts+string(os.PathSeparator)+namef))
		}
	} else {
		p = append(p, Post{Id: "ПОСТОВ НЕТ", Title: "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)", ContentText: ""})
	}

	rr.HTML(200, "index", &PagePost{TitlePage: "Блог проектов kaefik", Posts: p})
}

// посты блога
func PostsHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
}

// посты блога
func HtmlHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	namefs := Getlistfileindirectory(pathhtmlts)
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

func main() {
	m := martini.Classic()

	//	if !parse_args() {
	//		return
	//	}

	pathposts = "posts"
	pathhtml = "html"
	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Extensions: []string{".tmpl", ".html"}}))

	//	m.Use(auth.BasicFunc(authFunc))

	m.Get("/", indexHandler)
	m.Post("/html/:namepage", HtmlHandler)
	//	m.Get("/posts"--как это было --может и ничего и не было., PostsHandler)
	//	m.Post("/exec/:shop/:nstr", ExecHandler)
	m.RunOnAddr(":1111")

}

//func authFunc(username, password string) bool {
//	return (auth.SecureCompare(username, "admin") && auth.SecureCompare(password, "!!!!VVjhsdsajdbabjd1")) || (auth.SecureCompare(username, "mars") && auth.SecureCompare(password, "Verbat1mert")) || (auth.SecureCompare(username, "oilnur") && auth.SecureCompare(password, "Verbat1mqwe"))
//}

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
