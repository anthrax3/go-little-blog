// Posts
package models

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
// t-
func (p *Post) Print() {
	fmt.Println("Id post: 			", p.Id)
	fmt.Println("Date create post:  ", p.Date)
	fmt.Println("Draft create post:  ", p.Draft)
	fmt.Println("Title post: 		", p.Title)
	fmt.Println("SmallContentText post: 	", p.SmallContentText)
	fmt.Println("ContentText post: 	", p.ContentText)
}

// новый пост, создается как черновик
func (p *Post) New() {
	p.Id = ""
	p.SetDraft(true) // черновик
	p.Title = ""
	p.SmallContentText = ""
	p.ContentText = ""
	p.Date = utils.GetNowDate()

}

//сохранить пост в файл
// t-
func (p *Post) SavetoFile(namef string) {
	p.New()
	stitle := "title: " + "\"" + p.Title + "\"" + "\n"
	sdate := "date: " + "\"" + p.Date + "\"" + "\n"
	sdraft := "draft: " + "\"" + utils.Bool2String(p.Draft) + "\"" + "\n"
	str := beginTitlePost + "\n" + sdate + stitle + sdraft + endTitlePost + "\n" + p.ContentText + "\n"
	utils.Savestrtofile(namef, str)
}

//сохранить пост в файл c уникальным номером в имени файла, возвращает имени файла который создался
//  t?
func (p *Post) SavetoUniqFile(pathposts string) string {
	var uniqname string
	namefs := utils.Getlistfileindirectory(pathposts)
	namefs = utils.SorttoDown(namefs)
	kolfiles := len(namefs)
	if kolfiles == 0 {
		uniqname = "0"
	} else {
		numuniq, err := strconv.Atoi(utils.SplitFileName(namefs[0]))
		if err != nil {
			fmt.Println("Error in func SavetoUniqFile ", err)
			uniqname = "0"
		} else {
			uniqname = strconv.Itoa(numuniq + 1)
		}
	}
	uniqname = uniqname + ".md"
	p.SavetoFile(pathposts + string(os.PathSeparator) + uniqname)
	return uniqname
}

// полчение текста поста блога из файла : первая строка это заголовок сообщения, вторая и последующие это само сообщение - возвращает кол-во черновых сообщений (draft=true)
// t+
func (p *Post) GetPostfromFileMd(namef string) {
	var (
		titleRegexp   = regexp.MustCompile(`title:\s*\".+\"`)
		dateRegexp    = regexp.MustCompile(`date:\s*\".+\"`)
		draftRegexp   = regexp.MustCompile(`draft:\s*\".+\"`)
		contentRegexp = regexp.MustCompile(`\".+\"`)
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

// возвращает сообщение не черновик, если нет нормальных сообщений то возращается -1, иначе возвращается текущий номер позиции
//t+
func (p *Post) GetNormalPost(namefs []string, npos int) (int, int) {
	var koldraft int
	if (npos < 0) || (npos >= len(namefs)) {
		return -1, 0
	}
	koldraft = 0
	for npos < len(namefs) {
		p.GetPostfromFileMd(namefs[npos])
		if p.GetDraft() {
			npos += 1
			koldraft += 1
		} else {
			break
		}
	}
	if p.GetDraft() {
		return -1, 0
	} else {
		return npos, koldraft
	}
}

// получить сообщения из папки pathposts в кол-ве kolpost начиная с позиции tekpos - возвр-ет массив  []Post,кол-во файлов в папке, кол-во черновых сообщений
// t+
func GetPostsNewPos(pathposts string, tekpos int, kolpost int) ([]Post, int, int) {
	var (
		pp       Post
		kolfiles int
		koldraft int = 0
		kd       int
	)

	p := make([]Post, 0)
	koldraft = 0
	namefs := utils.Getlistfileindirectory(pathposts)
	namefs = utils.SorttoDown(namefs)
	namefs = utils.ConcatPathFileName(namefs, pathposts+string(os.PathSeparator))
	kolfiles = len(namefs)
	if tekpos < 0 {
		tekpos = 0
	}
	if !(tekpos < kolfiles) {
		tekpos = kolfiles - 1
	}

	tekkolpost := 0 // кол-во постов
	for (tekpos != -1) && (tekpos < kolfiles) && (tekkolpost < kolpost) {
		tekpos, kd = pp.GetNormalPost(namefs, tekpos)
		if tekpos != -1 {
			p = append(p, pp)
			tekpos += 1
			tekkolpost += 1
			koldraft += kd
		}
	}

	return p, kolfiles, koldraft
}

//------------ END методы структуры Post
