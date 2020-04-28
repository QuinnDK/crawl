package parse

import (
	"crawl/engine"
	"regexp"
)

func ParseContent(content []byte) engine.ParseResult {
	//<a href="/tag/哲学" class="tag">哲学</a>
	//re:=regexp.MustCompile(`<a href="([^"]+)">([^</a>]+)</a>`)
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`) //外边的小点 ` 并非是引号，我靠，花了好长时间
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requesrts = append(result.Requesrts, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
	//for _, m := range matches {
	//	fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
	//
	//}
}
