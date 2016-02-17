// go-little-blog_test
package main

import (
	"fmt"
	//	"os"
	"testing"

	//	"go-little-blog/utils"
)

func TestParseCfgFile(t *testing.T) {
	fmt.Println(ParseCfgFile("config.cfg"))
}
