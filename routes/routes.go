// routes
package routes

import (
	//	"fmt"
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

//-----------END вспомогательная функция которую надо будет удалить со временем

func IndexHandler(rr render.Render, w http.ResponseWriter, r *http.Request) {
	var pp models.Post
	p := make([]models.Post, 0)
	namefs := utils.Getlistfileindirectory(Pathposts)
	tnamefs := namefs
	vsegopost := len(namefs)
	if len(namefs) != 0 {
		namefs = utils.SorttoDown(namefs)
		if Kolpost > len(namefs) {
			tnamefs = namefs[:]
		} else {
			tnamefs = namefs[:Kolpost]
		}

		for _, namef := range tnamefs {
			pp.GetPostfromFileMd(Pathposts + string(os.PathSeparator) + namef)
			p = append(p, pp)
		}
	} else {
		p = append(p, models.Post{Id: "ПОСТОВ НЕТ", Title: "ЭТОТ БЛОГ ПУСТ. ПРИХОДИТЕ ПОЗЖЕ ;)", ContentText: ""})
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
