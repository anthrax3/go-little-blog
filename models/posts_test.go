// models_test
package models

import (
	//	"fmt"
	"os"
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
	res := p.GetNormalPost(namefs, 0)
	if p.Id != pathtestpostmd+"6.md" {
		t.Fatalf("должен быть p.Id - "+pathtestpostmd+"6.md", p.Id)
	}
	if res != 1 {
		t.Fatalf("должен быть res == 1  - ", res)
	}
	//----
	p.New()
	res = p.GetNormalPost(namefs, -1)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}
	//----
	p.New()
	res = p.GetNormalPost(namefsnull, 0)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}
	//----
	p.New()
	res = p.GetNormalPost(namefs, 1)
	if p.Id != pathtestpostmd+"6.md" {
		t.Fatalf("должен быть p.Id - "+pathtestpostmd+"6.md", p.Id)
	}
	if res != 1 {
		t.Fatalf("должен быть res == 1  - ", res)
	}
	//----
	p.New()
	res = p.GetNormalPost(namefs, 3)
	if p.Id != "" {
		t.Fatalf("должен быть p.Id пустой - ", p.Id)
	}
	if res != -1 {
		t.Fatalf("должен быть res == -1  - ", res)
	}

}
