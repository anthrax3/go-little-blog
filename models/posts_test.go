// models_test
package models

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"go-little-blog/utils"
)

func TestGetNormalPost(t *testing.T) {
	var (
		// 5 и 7 - черновики
		namefs                = []string{"7.md", "6.md", "5.md"}
		namefsnull            = []string{}
		pathtestpostmd string = "test-postmd" + string(os.PathSeparator) // путь где нах-ся тестовые файлы сообщений в формате маркдоуна
		p              Post
	)

	namefs = utils.ConcatPathFileName(namefs, pathtestpostmd)
	namefsnull = utils.ConcatPathFileName(namefsnull, pathtestpostmd)
	//----
	p.New()
	res, koldraft := p.GetNormalPost(namefs, 0)
	fmt.Println("koldraft", koldraft)
	if p.Id != pathtestpostmd+"6.md" {
		t.Fatalf("должен быть p.Id - "+pathtestpostmd+"6.md", p.Id)
	}
	if res != 1 {
		t.Fatalf("должен быть res == 1  - ", res)
	}
	if koldraft != 1 {
		t.Fatalf("неправильный результат  - ", koldraft)
	}

	//----
	p.New()
	res, koldraft = p.GetNormalPost(namefs, -1)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}
	//----
	p.New()
	res, koldraft = p.GetNormalPost(namefsnull, 0)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}
	//----
	p.New()
	res, koldraft = p.GetNormalPost(namefs, 1)
	if p.Id != pathtestpostmd+"6.md" {
		t.Fatalf("должен быть p.Id - "+pathtestpostmd+"6.md", p.Id)
	}
	if res != 1 {
		t.Fatalf("должен быть res == 1  - ", res)
	}
	//----
	p.New()
	res, koldraft = p.GetNormalPost(namefs, 3)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}

}

func TestGetPostsNewPos(t *testing.T) {
	var (
		pathposts string = "test-postmd"
		kolpost   int    = 3
	)

	resposts, _, _ := GetPostsNewPos(pathposts, 0, kolpost)

	//	if kolfiles != 8 {
	//		t.Fatalf("должен быть kolfiles == 8  - ", kolfiles)
	//	}
	if (resposts[0].Id != pathposts+string(os.PathSeparator)+"6.md") && (resposts[1].Id != pathposts+string(os.PathSeparator)+"5.md") && (resposts[2].Id != pathposts+string(os.PathSeparator)+"4.md") {
		t.Fatalf("должен быть 6.md , 5.md, 4.md  - ", resposts[0].Id+" "+resposts[1].Id+" "+resposts[2].Id+" ")
	}

	if (resposts[0].SmallContentText == "") && (resposts[1].SmallContentText == "") && (resposts[2].SmallContentText == "") {
		t.Fatalf("SmallContentText должен быть не пустой - ", resposts[0].SmallContentText+" -- "+resposts[1].SmallContentText+" -- "+resposts[2].SmallContentText+" -- ")
	}
	//----
	resposts, _, _ = GetPostsNewPos(pathposts, -1, kolpost)

	//	if kolfiles != 8 {
	//		t.Fatalf("должен быть kolfiles == 8  - ", kolfiles)
	//	}
	if (resposts[0].Id != pathposts+string(os.PathSeparator)+"6.md") && (resposts[1].Id != pathposts+string(os.PathSeparator)+"5.md") && (resposts[2].Id != pathposts+string(os.PathSeparator)+"4.md") {
		t.Fatalf("должен быть 6.md , 5.md, 4.md  - ", resposts[0].Id+" "+resposts[1].Id+" "+resposts[2].Id+" ")
	}

	if (resposts[0].SmallContentText == "") && (resposts[1].SmallContentText == "") && (resposts[2].SmallContentText == "") {
		t.Fatalf("SmallContentText должен быть не пустой - ", resposts[0].SmallContentText+" -- "+resposts[1].SmallContentText+" -- "+resposts[2].SmallContentText+" -- ")
	}
}

func TestSavetoUniqFile(t *testing.T) {
	//	var p Post
	//	p.New()
	//	p.SavetoUniqFile("test-postmd")

}

//func (p *Post) GetPostfromFileMd(namef string) {
func TestGetPostfromFileMd(t *testing.T) {
	var (
		p         Post
		pathposts string = "test-postmd"
	)

	p.GetPostfromFileMd(pathposts + string(os.PathSeparator) + "4.md")

	if strings.Compare(p.ContentText, utils.ConvertMarkdownToHtml("for test posts")) != 0 {
		p.Print()
		t.Fatalf("неправильный результат")
	}
	if p.Draft {
		t.Fatalf("неправильный результат", p.Draft)
	}
	if strings.Compare(p.Date, "2016-02-01") != 0 {
		t.Fatalf("неправильный результат", p.Date)
	}
	//
	if strings.Compare(p.Title, "4 posts Темы нету.") != 0 {
		t.Fatalf("неправильный результат", p.Title)
	}
	if strings.Compare(p.Id, pathposts+string(os.PathSeparator)+"4.md") != 0 {
		t.Fatalf("неправильный результат", p.Title)
	}
	// проверка на открытие несуществующего файла
	p.GetPostfromFileMd(pathposts + string(os.PathSeparator) + "1114.md")
	if strings.Compare(p.ContentText, utils.ConvertMarkdownToHtml("")) != 0 {
		p.Print()
		t.Fatalf("неправильный результат")
	}
	if p.Draft {
		t.Fatalf("неправильный результат", p.Draft)
	}
	if strings.Compare(p.Date, "") != 0 {
		t.Fatalf("неправильный результат", p.Date)
	}
	//
	if strings.Compare(p.Title, "") != 0 {
		t.Fatalf("неправильный результат", p.Title)
	}
	if strings.Compare(p.Id, pathposts+string(os.PathSeparator)+"1114.md") != 0 {
		t.Fatalf("неправильный результат", p.Title)
	}
}

//сохранить пост в файл
//func (p *Post) SavetoFile(namef string) {
func TestSavetoFile(t *testing.T) {
}
