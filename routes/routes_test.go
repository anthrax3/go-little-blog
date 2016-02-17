// routes_test
package routes

import (
	//	"fmt"
	"os"
	"testing"

	//	"go-little-blog/utils"
)

// получить сообщения из папки pathposts в кол-ве kolpost
//func GetPostsNew(pathposts string, kolpost int) ([]models.Post, int) {

func TestGetPostsNew(t *testing.T) {
	var (
		pathposts string = "test-postmd"
		kolpost   int    = 3
	)

	resposts, kolfiles := GetPostsNew(pathposts, kolpost)

	if kolfiles != 8 {
		t.Fatalf("должен быть kolfiles == 8  - ", kolfiles)
	}
	if (resposts[0].Id != pathposts+string(os.PathSeparator)+"6.md") && (resposts[1].Id != pathposts+string(os.PathSeparator)+"5.md") && (resposts[2].Id != pathposts+string(os.PathSeparator)+"4.md") {
		t.Fatalf("должен быть 6.md , 5.md, 4.md  - ", resposts[0].Id+" "+resposts[1].Id+" "+resposts[2].Id+" ")
	}

	if (resposts[0].SmallContentText == "") && (resposts[1].SmallContentText == "") && (resposts[2].SmallContentText == "") {
		t.Fatalf("SmallContentText должен быть не пустой - ", resposts[0].SmallContentText+" -- "+resposts[1].SmallContentText+" -- "+resposts[2].SmallContentText+" -- ")
	}

}
