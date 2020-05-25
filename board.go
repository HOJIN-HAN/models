package models // import "github.com/HOJIN-HAN/app"

type Board struct {
	Id    int64 `db:"post_id"`
	Title string
	Body  string
	Nick  string
}
