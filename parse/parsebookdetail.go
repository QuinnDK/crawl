package parse

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
	"strconv"
)

var title = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publuc = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+ )</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p>[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte, bookname string) engine.ParseResult {

	bookedetail := model.Bookdetail{}

	bookedetail.Title = ExtraString(contents, title)
	bookedetail.Author = ExtraString(contents, autoRe)
	bookedetail.Public = ExtraString(contents, publuc)
	page, err := strconv.Atoi(ExtraString(contents, pageRe))
	if err == nil {
		bookedetail.Booktages = page
	}
	bookedetail.Price = ExtraString(contents, priceRe)
	bookedetail.Score = ExtraString(contents, scoreRe)
	bookedetail.Into = ExtraString(contents, intoRe)

	result := engine.ParseResult{
		Items: []interface{}{bookedetail},
	}

	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

	return ""
}
