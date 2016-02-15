// Posts
package models

import (
	"fmt"
	"regexp"
	"strings"

	"go-little-blog/utils"
)

// структура поста в блоге
type Post struct {
	Id               string
	Draft            bool   // true - черновик сообщения
	Date             string // дата создания поста
	Title            string
	SmallContentText string // часть сообщения ContentText
	ContentText      string // весь сообщение блога
}

var (
	beginTitlePost string = "---"
	endTitlePost   string = "---"
)

//------------ методы структуры Post

func (p *Post) SetDraft(b bool) {
	p.Draft = b
}

func (p *Post) GetDraft() bool {
	return p.Draft
}

// вывод на экран
func (p *Post) Print() {
	fmt.Println("Id post: 			", p.Id)
	fmt.Println("Date create post:  ", p.Date)
	fmt.Println("Draft create post:  ", p.Draft)
	fmt.Println("Title post: 		", p.Title)
	fmt.Println("SmallContentText post: 	", p.SmallContentText)
	fmt.Println("ContentText post: 	", p.ContentText)
}

// новый пост
func (p *Post) New() {
	p.Id = ""
	p.SetDraft(true) // черновик
	p.Title = "title newpost"
	p.SmallContentText = "small content new post"
	p.ContentText = "content new post"
	p.Date = utils.GetNowDate()

}

//сохранить пост в файл   - ПЕРЕДЕЛАТЬ, ДОБАВИТЬ ШАПКУ ПАРАМЕТРОВ draft
func (p *Post) SavetoFile(namef string) {
	p.New()
	stitle := "title: " + "\"" + p.Title + "\"" + "\n"
	sdate := "date: " + "\"" + p.Date + "\"" + "\n"
	sdraft := "draft: " + "\"" + utils.Bool2String(p.Draft) + "\"" + "\n"
	str := beginTitlePost + "\n" + sdate + stitle + sdraft + endTitlePost + "\n" + p.ContentText + "\n"
	utils.Savestrtofile(namef, str)
}

// полчение текста поста блога из файла : первая строка это заголовок сообщения, вторая и последующие это само сообщение
func (p *Post) GetPostfromFileMd(namef string) {
	var (
		titleRegexp   = regexp.MustCompile(`title:\s*\".+\"`)
		dateRegexp    = regexp.MustCompile(`date:\s*\".+\"`)
		draftRegexp   = regexp.MustCompile(`draft:\s*\".+\"`)
		contentRegexp = regexp.MustCompile(`\".+\"`)
		//		descRegexp   = regexp.MustCompile(`description:\s*\".+\"`)
	)
	stitle := ""
	scontent := ""
	sdraft := ""
	smallcontent := ""
	stitledate := ""
	stitlepost := make([]string, 0)
	pospost := -1
	str := utils.Readfiletxt(namef)
	lenstr := len(str)

	linestr := strings.Split(str, "\n")

	// выделение заголовка (описания) поста
	if lenstr > 2 {
		if (strings.Index(linestr[0], beginTitlePost)) != -1 {
			for i := 1; i < len(linestr); i++ {
				if (strings.Index(linestr[i], endTitlePost)) != -1 {
					pospost = i
					break
				} else {
					stitlepost = append(stitlepost, linestr[i])
				}

			}
		}
	}

	stitle = strings.Join(stitlepost, "\n")

	// выделение параметра date
	stitledate = dateRegexp.FindString(stitle)
	stitledate = contentRegexp.FindString(stitledate)
	if len(stitledate) != 0 {
		stitledate = stitledate[1 : len(stitledate)-1]
	}

	// выделение параметра draft
	sdraft = draftRegexp.FindString(stitle)
	sdraft = contentRegexp.FindString(sdraft)
	if len(sdraft) != 0 {
		sdraft = sdraft[1 : len(sdraft)-1]
	}

	// выделение параметра title
	stitle = titleRegexp.FindString(stitle)
	stitle = contentRegexp.FindString(stitle)
	if len(stitle) != 0 {
		stitle = stitle[1 : len(stitle)-1]
	}

	if (pospost+1 <= len(linestr)) && (pospost != -1) {
		scontent = strings.Join(linestr[pospost+1:], "\n")
	}

	//отсекаем 140 символами для краткого содержания поста
	rscontent := []rune(scontent)
	if (len(rscontent) != 0) && (len(rscontent) > 140) {
		smallcontent = string(rscontent[0:140])
	} else {
		smallcontent = scontent
	}

	*p = Post{Id: namef, Title: stitle, ContentText: utils.ConvertMarkdownToHtml(scontent), SmallContentText: utils.ConvertMarkdownToHtml(smallcontent), Date: stitledate, Draft: utils.String2Bool(sdraft)}
}

//------------ END методы структуры Post
