// routes
package routes

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	//	"strings"

	"go-little-blog/models"
	"go-little-blog/utils"

	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
)

var (
	Pathposts    string // папка в которой нах-ся посты блога
	Pathhtml     string // папка в которой нах-ся обычные html страницы блога
	Pathtemplate string //  папка в которой нах-ся шаблоны
	Kolpost      int    // кол-во постов (сообщений) на главной странице блога
	Tekpost      int    // номер сообщения с которого начинается сообщения на странице
)

// получить сообщения из папки pathposts в кол-ве kolpost
func GetPostsNew(pathposts string, kolpost int) ([]models.Post, int) {
	var (
		pp       models.Post
		kolfiles int
	)

	p := make([]models.Post, 0)

	namefs := utils.Getlistfileindirectory(pathposts)
	namefs = utils.SorttoDown(namefs)
	namefs = utils.ConcatPathFileName(namefs, pathposts+string(os.PathSeparator))
	kolfiles = len(namefs)

	tekkolpost := 0 // кол-во постов
	tekpos := 0
	for (tekpos != -1) && (tekpos < kolfiles) && (tekkolpost < kolpost) {
		tekpos = pp.GetNormalPost(namefs, tekpos)
		if tekpos != -1 {
			p = append(p, pp)
			tekpos += 1
			tekkolpost += 1
		}
	}

	return p, kolfiles
}

// получить сообщения из папки pathposts в кол-ве kolpost
func getPosts(pathposts string, kolpost int) ([]models.Post, int) {
	var (
		//		tnamefs []string
		//		pp       models.Post
		kolfiles int
	)

	p := make([]models.Post, 0)

	namefs := utils.Getlistfileindirectory(pathposts)
	namefs = utils.SorttoDown(namefs)
	kolfiles = len(namefs)

	//	if (kolfiles != 0) && (kolpost != 0) {

	//		if kolpost > len(namefs) {
	//			tnamefs = namefs[:]
	//			for npos, _ := range tnamefs {
	//				if pp.GetNormalPost(tnamefs, npos) {
	//					p = append(p, pp)
	//				}
	//			}

	//		} else {
	//			tnamefs = namefs[:kolpost]
	//		}

	//		for _, namef := range tnamefs {
	//			pp.GetPostfromFileMd(Pathposts + string(os.PathSeparator) + namef)
	//			if !pp.GetDraft() { // не отражаются черновики
	//				p = append(p, pp)
	//			} else {
	//				kolpost += 1
	//				if kolpost > len(namefs) {
	//				}
	//			}

	//		}
	//	} else {
	//		pp.New()
	//		pp.Title = "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)"
	//		p = append(p, pp)
	//	}
	return p, kolfiles
}

//-----------END вспомогательная функция которую надо будет удалить со временем

func IndexHandler(rr render.Render, w http.ResponseWriter, r *http.Request) {
	//	var pp models.Post
	//	p := make([]models.Post, 0)
	fmt.Println(Pathposts)
	p, vsegopost := GetPostsNew(Pathposts, Kolpost)
	fmt.Println(len(p))
	for _, v := range p {
		fmt.Print(v)
		//		fmt.Print(v.Id)
		//		fmt.Print("  ", v.Draft)
		//		fmt.Println("  ", v.SmallContentText)

	}
	rr.HTML(200, "index", &models.PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postright: vsegopost - Kolpost})
}

// посты блога
func PostsHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
}

// посты блога
func HtmlHandler(rr render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	namefs := utils.Getlistfileindirectory(Pathhtml)
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
	var pp models.Post
	p := make([]models.Post, 0)
	numpost, _ := strconv.Atoi(params["numpost"])
	namefs := utils.Getlistfileindirectory(Pathposts)
	namefs = utils.SorttoUp(namefs)
	tnamefs := namefs
	vsegopost := len(namefs)

	if numpost <= 0 {
		numpost = len(namefs)
	}

	if numpost >= len(namefs) {
		if Kolpost > len(namefs) {
			tnamefs = namefs[:]
		} else {
			tnamefs = namefs[len(namefs)-Kolpost:]
		}

		if len(namefs) != 0 {
			for k := len(tnamefs) - 1; k >= 0; k-- {
				namef := tnamefs[k]
				pp.GetPostfromFileMd(Pathposts + string(os.PathSeparator) + namef)
				p = append(p, pp)
			}
		} else {
			p = append(p, models.Post{Id: "ПОСТОВ НЕТ", Title: "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)", ContentText: ""})
		}
		rr.HTML(200, "view", &models.PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postright: vsegopost - Kolpost})
		return
	}
	if len(namefs) != 0 {
		kk := numpost - Kolpost
		if kk <= 0 {
			kk = 0
		} else {
			kk = numpost - Kolpost // - 1
		}
		tnamefs := namefs[kk:numpost]

		if len(namefs) != 0 {
			for k := len(tnamefs) - 1; k >= 0; k-- {
				namef := tnamefs[k]
				pp.GetPostfromFileMd(Pathposts + string(os.PathSeparator) + namef)
				p = append(p, pp)
			}
		}
	} else {
		rr.Redirect("/")
		return
	}
	rr.HTML(200, "view", &models.PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postleft: numpost + Kolpost, Postright: numpost - Kolpost})
}
