package main

import (
	//	"flag"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	"go-little-blog/models"
	"go-little-blog/routes"
	"go-little-blog/utils"

	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
)

////------------ Объявление типов и глобальных переменных

// вывод на печать массива Post
func Print(p []models.Post) {
	for _, v := range p {
		v.Print()
	}
}

//------------ END Объявление типов и глобальных переменных

func parseCmdArgs() bool {
	res := false
	var p models.Post
	lenargs := len(os.Args)
	if lenargs == 2 {
		if os.Args[1] == "help" {
			fmt.Println("new post -  create new post")
			res = true
		}
	}
	if lenargs == 3 { // два аргумента
		if (os.Args[1] == "new") && (os.Args[2] == "post") {

			p.New()
			//			p.SavetoFile(routes.Pathposts + string(os.PathSeparator) + "newfile.md")
			fmt.Println("Result create new post : ", p.SavetoUniqFile(routes.Pathposts))
			res = true
		}
	}
	return res
}

// возвращает значение параметра params из строки str
func GetParamsFromStr(params string, str string) string {
	var val string = ""
	pos := strings.Index(str, params+":")
	if (pos == -1) && (len(params) == 0) {
		return ""
	}
	val = strings.TrimLeft(strings.TrimRight(str[pos+len(params)+1:], " "), " ")
	return val

}

//парсинг конфиг файла map[ключ] значение_ключа
func ParseCfgFile(namef string) map[string]string {
	var (
		pathpostsRegexp = regexp.MustCompile(`pathposts:.+`)
		//		contentRegexp   = regexp.MustCompile(`\".+\"`)
	)
	res := make(map[string]string, 0)
	str := utils.Readfiletxt(namef)

	if len(str) == 0 {
		return res
	}

	pathposts := pathpostsRegexp.FindString(str)
	res["pathposts"] = pathposts

	return res
}

func main() {
	fmt.Println("Start...")
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
	routes.Pathposts = "posts"
	routes.Pathhtml = "html"
	//	routes.Pathtemplate = "templates"
	routes.Pathtemplate = "templates" + string(os.PathSeparator) + "uno-theme"
	routes.Kolpost = 3 // кол-во постов которые видны на странице
	//--------------

	if parseCmdArgs() {
		return
	} else {

		unescapeFuncMap := template.FuncMap{"unescape": utils.Unescape}

		staticOptions := martini.StaticOptions{Prefix: "assets"}
		m.Use(martini.Static("assets", staticOptions))

		m.Use(render.Renderer(render.Options{
			Directory:  routes.Pathtemplate,                 // Specify what path to load the templates from.
			Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
			Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
			IndentJSON: true,                                // Output human readable JSON
			Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
			Extensions: []string{".tmpl", ".html"}}))

		m.Get("/", routes.IndexHandler)
		m.Get("/html/:namepage", routes.HtmlHandler)
		m.Get("/view/:numpost", routes.ViewHandler)
		m.RunOnAddr(":" + parports)
	}
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
