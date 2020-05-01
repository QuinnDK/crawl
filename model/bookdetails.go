package model

import "strconv"

type Bookdetail struct {
	Title     string
	Author    string
	Public    string
	Booktages int
	Price     string
	Score     string
	Into      string
}

func (b Bookdetail) String() string {
	return "书名: " + b.Title + " 作者: " + b.Author + " 出版社:" + b.Public + " 页数: " + strconv.Itoa(b.Booktages) + "  价格: " + b.Price + " 评分:" + b.Score + " 内容简介:" + b.Into
}
