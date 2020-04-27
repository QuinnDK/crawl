package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	//resp,err:=http.Get("https://book.douban.com/top250?start=0")  豆瓣已经加入的发爬虫机制，这种方法不能用了
	res, err := http.NewRequest("GET", "https://book.douban.com", nil)

	if err != nil {
		panic(err)
	}
	res.Header.Add("User-Agent", "test")

	client := &http.Client{}
	resp, err2 := client.Do(res)

	if err2 != nil {
		err = err2
		return
	}
	defer resp.Body.Close() //延迟关闭

	if resp.StatusCode != http.StatusOK { //statusOK 值为两百，像404等等为服务器异常
		fmt.Printf("Error status code : %d", resp.StatusCode)

	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", result)
	ParseContent(result)

	//正则表达式
	//	str:="asdajidasfksaiaioaskfivsvs"
	//	re:=regexp.MustCompile(".*ksai")
	//	rest:=re.FindString(str)
	//	fmt.Println(rest)
}

func ParseContent(content []byte) {
	//<a href="/tag/哲学" class="tag">哲学</a>
	//re:=regexp.MustCompile(`<a href="([^"]+)">([^</a>]+)</a>`)
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`) //外边的小点 ` 并非是引号，我靠，花了好长时间
	matches := re.FindAllSubmatch(content, -1)

	for _, m := range matches {
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))

	}

}
