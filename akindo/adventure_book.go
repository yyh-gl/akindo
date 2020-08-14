package akindo

import (
	"io"
	"log"
	"os"
)

// adventureBookWriter : 冒険の書著者を表す構造体
type adventureBookWriter struct {
	book *log.Logger
}

// newAdventureBookWriter : 冒険の書著者を召喚
func newAdventureBookWriter() adventureBookWriter {
	var w io.Writer
	switch os.Getenv("ENV") {
	default:
		w = os.Stdout
	}
	return adventureBookWriter{
		book: log.New(w, "[Akindo] ", log.LstdFlags),
	}
}

// WriteTradeLog : 取引記録に書き込む
func (abw adventureBookWriter) WriteTradeLog(s string) {
	abw.book.Println(s)
}

// WriteAccident : 盗賊に襲われた記録を書き込む
func (abw adventureBookWriter) WriteAccident(err error) {
	abw.book.Fatal(err)
}
