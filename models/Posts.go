// Posts
package models

import (
	"fmt"
	"strings"

	"go-little-blog/utils"
)

// структура поста в блоге
type Post struct {
	Id          string
	Title       string
	ContentText string
}

//------------ методы структуры Post

// вывод на экран
func (p *Post) Print() {
	fmt.Println("Id post: ", p.Id)
	fmt.Println("Title post: ", p.Title)
	fmt.Println("ContentText post: ", p.ContentText)
}

// новый пост
func (p *Post) New() {
	p.Title = "новое сообщение"
	p.ContentText = "пустое сообщение"
}

//сохранить пост в файл
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
	p = &Post{Id: namef, Title: utils.ConvertMarkdownToHtml(stitle), ContentText: utils.ConvertMarkdownToHtml(scontent)}
	//	res := Post{Id: namef, Title: stitle, ContentText: scontent}
	//	return res
}

//------------ END методы структуры Post
