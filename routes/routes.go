// routes
package routes

import (
	//	"fmt"
	"net/http"
	"strconv"

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
	p, vsegopost := models.GetPostsNewPos(Pathposts, 0, Kolpost)
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
	//	var pp models.Post
	//	p := make([]models.Post, 0)
	numpost, _ := strconv.Atoi(params["numpost"])

	p, kolfiles := models.GetPostsNewPos(Pathposts, numpost, Kolpost)

	if kolfiles == 0 {
		rr.Redirect("/")
		return

	}

	rr.HTML(200, "view", &models.PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postleft: numpost + Kolpost - 1, Postright: numpost - Kolpost + 1})
	//	rr.HTML(200, "view", &models.PagePost{TitlePage: "Блог проектов kaefik", Posts: p, Postleft: numpost + 1, Postright: numpost - 1})
}
