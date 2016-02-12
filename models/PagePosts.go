// PagePosts
package models

//данные для генерации страницы html
type PagePost struct {
	TitlePage string
	Posts     []Post
	Postleft  int // кол-во сообщений влево , т.е. более поздние
	Postright int // кол-во сообщений вправо , т.е. более ранние
}
