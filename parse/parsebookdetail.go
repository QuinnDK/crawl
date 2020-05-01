package parse

import (
	"regexp"
)

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[/d/D]*?<a.*?>([^<]+ )</a>`)
var publuc = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+ )<br>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+ )<br>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+ )<br>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+ )</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[/d/D]*?<p>([^<]+ )</p></div>`)

func ParseBookDetail(contents []byte) {

	//bookedetail:=model.Bookdetail{}
	//match:=autoRe.FindAllSubmatch(contents,-1)

}
