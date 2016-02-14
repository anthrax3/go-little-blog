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
	Id          string
	Date        string // дата создания поста
	Title       string
	ContentText string
}

var (
	beginTitlePost string = "---"
	endTitlePost   string = "---"
)

//------------ методы структуры Post

// вывод на экран
func (p *Post) Print() {
	fmt.Println("Id post: 			", p.Id)
	fmt.Println("Date create post:  ", p.Date)
	fmt.Println("Title post: 		", p.Title)
	fmt.Println("ContentText post: 	", p.ContentText)
}

// новый пост
func (p *Post) New() {
	p.Title = "новое сообщение"
	p.ContentText = "пустое сообщение"
}

//сохранить пост в файл   - ПЕРЕДЕЛАТЬ, ДОБАВИТЬ ШАПКУ ПАРАМЕТРОВ title и  date
func (p *Post) SavetoFile(namef string) {
	p.Title = "новое сообщение"
	p.ContentText = "пустое сообщение"
	str := p.Title + "\n" + p.ContentText + "\n"
	utils.Savestrtofile(namef, str)
}

// полчение текста поста блога из файла : первая строка это заголовок сообщения, вторая и последующие это само сообщение
func (p *Post) GetPostfromFile(namef string) {
	str := utils.Readfiletxt(namef)
	sposts := strings.Index(str, "\n") // поиск первой строки - заголовка сообщения
	stitle := ""
	scontent := ""
	if sposts != -1 {
		stitle = str[0:sposts]
		if sposts+1 <= len(str) {
			scontent = str[sposts+1:]
		}
	}
	*p = Post{Id: namef, Title: utils.ConvertMarkdownToHtml(stitle), ContentText: utils.ConvertMarkdownToHtml(scontent)}
}

// полчение текста поста блога из файла : первая строка это заголовок сообщения, вторая и последующие это само сообщение
func (p *Post) GetPostfromFileMd(namef string) {
	var (
		titleRegexp   = regexp.MustCompile(`title:\s*\".+\"`)
		dateRegexp    = regexp.MustCompile(`date:\s*\".+\"`)
		contentRegexp = regexp.MustCompile(`\".+\"`)
		//		descRegexp   = regexp.MustCompile(`description:\s*\".+\"`)
	)
	stitle := ""
	scontent := ""
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

	stitledate = dateRegexp.FindString(stitle) // выделение параметра date
	stitledate = contentRegexp.FindString(stitledate)
	if len(stitledate) != 0 {
		stitledate = stitledate[1 : len(stitledate)-1]
	}
	stitle = titleRegexp.FindString(stitle) // выделение параметра title
	stitle = contentRegexp.FindString(stitle)
	if len(stitle) != 0 {
		stitle = stitle[1 : len(stitle)-1]
	}

	if (pospost+1 <= len(linestr)) && (pospost != -1) {
		for j := pospost + 1; j < len(linestr); j++ {
			scontent += linestr[j] + "\n"
		}
	}

	*p = Post{Id: namef, Title: stitle, ContentText: utils.ConvertMarkdownToHtml(scontent), Date: stitledate}
}

//------------ END методы структуры Post
